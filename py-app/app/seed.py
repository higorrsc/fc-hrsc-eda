from app.database import SessionLocal, init_db
from app.kafka_producer import send_transaction
from app.models import AccountBalance
from sqlalchemy.orm import Session


def create_seed_data(db_session: Session):
    """
    Populate the database with initial seed data for account balances.

    Args:
        db_session (Session): The database session used to add seed data.

    The function creates a list of account balance data with randomly
    generated account IDs and predefined balances. These records are
    added to the database and committed. Prints a confirmation message
    once seed data has been successfully added.
    """
    initial_data = [
        {
            "account_id": "77bf4d6a-e7a5-40a5-b131-0b906c2e4d9e",
            "balance": 1000.0,
        },
        {
            "account_id": "d33e99a3-27aa-4891-a93d-220204013a55",
            "balance": 2500.0,
        },
        {
            "account_id": "a1ca32d6-f939-47bf-8644-09b6949ba750",
            "balance": 500.0,
        },
        {
            "account_id": "00faaa0e-c2fd-429f-b8b6-eac3d0a9d414",
            "balance": 0.0,
        },
    ]
    try:
        existing_accounts = {
            account.account_id: account.balance
            for account in db_session.query(AccountBalance).all()
        }

        new_accounts = []

        for data in initial_data:
            account_id = data["account_id"]
            balance = data["balance"]

            if account_id not in existing_accounts:
                new_accounts.append(AccountBalance(**data))

            send_transaction(account_id, balance)

        if new_accounts:
            db_session.add_all(new_accounts)
            db_session.commit()
            print("✅ Seed data added successfully.")
        else:
            print("✔️ Seed data already exists.")
    except Exception as e:
        print(f"❌ Error adding seed data: {e}")
        db_session.rollback()
    finally:
        db_session.close()


if __name__ == "__main__":
    init_db()
    db = SessionLocal()
    create_seed_data(db)
