import json
import threading

from app.models import AccountBalance
from confluent_kafka import Consumer, KafkaException
from sqlalchemy.orm import Session


class KafkaConsumerService:
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
                "bootstrap.servers": "kafka:9092",
                "group.id": "python-microservice",
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
                if msg.error().code() != KafkaException._PARTITION_EOF:
                    print(f"Kafka error: {msg.error()}")
                continue

            # Processar a mensagem recebida
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
        data = json.loads(message)
        account_id = data["account_id"]
        amount = data["amount"]

        # Atualizar o saldo da conta no banco de dados
        db: Session = next(self.db_session())
        account = db.query(AccountBalance).filter_by(account_id=account_id).first()
        if not account:
            account = AccountBalance(account_id=account_id, balance=0.0)
            db.add(account)
        account.balance += amount
        db.commit()

    def stop(self):
        self.running = False
        self.consumer.close()
