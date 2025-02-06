from uuid import UUID

from app.database import SessionLocal, get_db, init_db
from app.kafka_consumer import KafkaConsumerService
from app.kafka_setup import create_kafka_topic, wait_for_kafka
from app.models import AccountBalance
from app.seed import create_seed_data
from fastapi import Depends, FastAPI, HTTPException
from sqlalchemy.orm import Session

app = FastAPI()


@app.get("/balances/{account_id}")
async def get_balance(account_id: UUID, db: Session = Depends(get_db)):
    """
    Get the balance of an account.

    Args:
        account_id (str): The id of the account.
        db (Session): The database session.

    Returns:
        dict: A dictionary with the account id and balance.

    Raises:
        HTTPException: If the account is not found.
    """
    print(f"Getting balance for account {account_id}")
    account = db.query(AccountBalance).filter_by(account_id=str(account_id)).first()
    if not account:
        raise HTTPException(status_code=404, detail="Account not found")
    return {"account_id": account_id, "balance": account.balance}


@app.on_event("startup")
async def startup_event():
    """
    Event handler for the application's startup event.

    This function initializes the database and starts the Kafka consumer service.

    Initializes:
        - The database using the `init_db` function.
        - The Kafka consumer service using the `KafkaConsumerService`.

    """

    init_db()

    db = SessionLocal()
    try:
        create_seed_data(db)
    finally:
        db.close()

    wait_for_kafka()
    create_kafka_topic()

    kafka_service = KafkaConsumerService(get_db)
    kafka_service.start()
