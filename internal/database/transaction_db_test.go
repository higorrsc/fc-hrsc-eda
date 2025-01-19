package database

import (
	"database/sql"
	"testing"

	"github.com/higorrsc/fc-hrsc-microservices/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	clientFrom    *entity.Client
	accountFrom   *entity.Account
	clientTo      *entity.Client
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (suite *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255) PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), created_at DATETIME, updated_at DATETIME)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255) PRIMARY KEY, client_id VARCHAR(255), balance FLOAT, created_at DATETIME, updated_at DATETIME)")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255) PRIMARY KEY, account_id_from VARCHAR(255), account_id_to VARCHAR(255), amount FLOAT, created_at DATETIME)")

	clientFrom, err := entity.NewClient("John Doe", "L4uQK@example.com")
	suite.Nil(err)
	suite.clientFrom = clientFrom

	clientTo, err := entity.NewClient("Jane Doe", "onpen@huitiho.lv")
	suite.Nil(err)
	suite.clientTo = clientTo

	accountFrom := entity.NewAccount(suite.clientFrom)
	accountFrom.Credit(1000.00)
	suite.accountFrom = accountFrom

	accountTo := entity.NewAccount(suite.clientTo)
	accountTo.Credit(1000.00)
	suite.accountTo = accountTo

	suite.transactionDB = NewTransactionDB(db)
}

func (suite *TransactionDBTestSuite) TearDownSuite() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE transactions")
	suite.db.Exec("DROP TABLE accounts")
	suite.db.Exec("DROP TABLE clients")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (suite *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(suite.accountFrom, suite.accountTo, 100.00)
	suite.Nil(err)
	err = suite.transactionDB.Create(transaction)
	suite.Nil(err)
}
