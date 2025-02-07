from sqlalchemy import create_engine
from sqlalchemy.orm import declarative_base, sessionmaker

DATABASE_URL = "mysql+pymysql://root:root@database-balance:3306/wallet"

engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()


def init_db():
    """
    Initializes the database, creating the tables.
    """
    Base.metadata.create_all(bind=engine)


def get_db():
    """
    Dependency to get a session from the database.

    Yields a database session. Automatically closes the session after the
    context is exited.
    """
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
