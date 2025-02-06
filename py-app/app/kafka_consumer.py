import json
import threading

from app.models import AccountBalance
from confluent_kafka import Consumer
from sqlalchemy.orm import Session


class KafkaConsumerService:
    """
    KafkaConsumerService is responsible for consuming messages from a Kafka topic.

    This service initializes a Kafka consumer and manages the consumption of
    messages in a separate thread. It processes each message by updating the
    account balance in the database.

    Attributes:
        db_session: A function to obtain a database session.
        running: A boolean flag indicating if the service is running.
        consumer: An instance of the Kafka Consumer.
    """

    def __init__(self, db_session):
        """
        Initialize a KafkaConsumerService.

        Args:
            db_session: A function that returns a session to the database.
        """

        self.db_session = db_session
        self.running = True
        self.consumer = Consumer(
            {
                "bootstrap.servers": "kafka:29092",
                "group.id": "wallet",
                "auto.offset.reset": "earliest",
            }
        )

    def start(self):
        """
        Start the Kafka consumer service.

        This method starts a new thread that consumes messages from the
        Kafka topic. This method is non-blocking.
        """
        threading.Thread(target=self._consume_messages, daemon=True).start()

    def _consume_messages(self):
        """
        Consume messages from the Kafka topic.

        This method runs in an infinite loop and will not return until the
        service is stopped. It will consume messages from the topic and
        process them by calling the _process_message method. If an error
        occurs while consuming messages, it will be printed to the console.
        """
        self.consumer.subscribe(["transactions"])

        while self.running:
            msg = self.consumer.poll(1.0)

            if msg is None:
                continue
            if msg.error():
                print(f"Kafka error: {msg.error()}")
                continue

            self._process_message(msg.value().decode("utf-8"))

    def _process_message(self, message):
        """
        Processes the received message from the Kafka transaction thread.

        The expected message will have the following format:
        {
            "account_id": string,
            "amount": float
        }

        The account balance will be updated in the database based on the
        received message.
        """
        try:
            data = json.loads(message)
            account_id = data["account_id"]
            amount = data["amount"]

            db: Session = next(self.db_session())
            account = db.query(AccountBalance).filter_by(account_id=account_id).first()
            if not account:
                account = AccountBalance(account_id=account_id, balance=0.0)
                db.add(account)
                account.balance += amount
                db.commit()
        except json.JSONDecodeError:
            print("Erro ao decodificar JSON")
        except Exception as e:
            print(f"Erro ao processar mensagem: {e}")

    def stop(self):
        """
        Stop the Kafka consumer service.

        This method stops the Kafka consumer service and closes the
        connection to the Kafka broker.
        """
        self.running = False
        self.consumer.close()
