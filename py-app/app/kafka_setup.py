import time

from confluent_kafka.admin import AdminClient, NewTopic

KAFKA_BOOTSTRAP_SERVERS = "kafka:29092"
TOPIC_NAME = "transactions"


def create_kafka_topic():
    """
    Create a Kafka topic if it does not already exist.
    """
    admin_client = AdminClient({"bootstrap.servers": KAFKA_BOOTSTRAP_SERVERS})

    topics_metadata = admin_client.list_topics(timeout=5)
    existing_topics = set(topic.topic for topic in topics_metadata.topics.values())

    if TOPIC_NAME not in existing_topics:
        print(f"🔹 Creating topic '{TOPIC_NAME}'...")
        new_topic = [NewTopic(TOPIC_NAME, num_partitions=1, replication_factor=1)]
        fs = admin_client.create_topics(new_topic)

        for topic, future in fs.items():
            try:
                future.result()
                print(f"✅ Topic '{topic}' created!")
            except Exception as e:
                print(f"❌ Failed to create topic '{topic}': {e}")
    else:
        print(f"✔️ Topic '{TOPIC_NAME}' already exists.")


def wait_for_kafka():
    """
    Wait for Kafka to be available before proceeding.
    """
    timeout = 30
    start_time = time.time()

    while time.time() - start_time < timeout:
        try:
            admin_client = AdminClient({"bootstrap.servers": KAFKA_BOOTSTRAP_SERVERS})
            if admin_client.list_topics(timeout=5).topics:
                print("✅ Kafka is available!")
                return
        except Exception:
            print("⏳ Waiting for Kafka...")
            time.sleep(5)

    raise TimeoutError("❌ Kafka is not available after 30 seconds.")


if __name__ == "__main__":
    time.sleep(5)
    create_kafka_topic()
