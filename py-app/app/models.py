import uuid

from app.database import Base
from sqlalchemy import Column, Float
from sqlalchemy.dialects.mysql import CHAR


class AccountBalance(Base):
    """
    Defines the table name and fields
    """

    __tablename__ = "account_balances"

    account_id = Column(CHAR(36), primary_key=True, default=lambda: str(uuid.uuid4()))
    balance = Column(Float, default=0.0)
