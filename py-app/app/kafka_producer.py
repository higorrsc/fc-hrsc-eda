import json
import time

from confluent_kafka import Producer

KAFKA_BOOTSTRAP_SERVERS = "kafka:29092"
TOPIC_NAME = "transactions"

producer_config = {"bootstrap.servers": KAFKA_BOOTSTRAP_SERVERS}
producer = Producer(producer_config)


def delivery_report(err, msg):
    """
    Callback called once for each message produced to indicate delivery status.
    """
    if err is not None:
        print(f"❌ Erro ao enviar mensagem: {err}")
    else:
        print(f"✅ Mensagem enviada para {msg.topic()} [{msg.partition()}]")


def send_transaction(account_id, amount):
    """
    Send a transaction message to the Kafka topic.
    """
    data = {
        "account_id": account_id,
        "amount": amount,
    }
    producer.produce(
        TOPIC_NAME,
        json.dumps(data).encode("utf-8"),
        callback=delivery_report,
    )
    producer.flush()


if __name__ == "__main__":
    print("🔹 Sending test messages...")
    for i in range(5):
        send_transaction(f"acc_{i}", 100 + i * 10)
        time.sleep(1)
