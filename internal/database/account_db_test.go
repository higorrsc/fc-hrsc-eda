package database

import (
	"database/sql"
	"testing"

	"github.com/higorrsc/fc-hrsc-microservices/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (suite *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255) PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), created_at DATETIME, updated_at DATETIME)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255) PRIMARY KEY, client_id VARCHAR(255), balance FLOAT, created_at DATETIME, updated_at DATETIME)")
	suite.accountDB = NewAccountDB(db)
	suite.client, _ = entity.NewClient("John Doe", "L4uQK@example.com")
}

func (suite *AccountDBTestSuite) TearDownSuite() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE accounts")
	suite.db.Exec("DROP TABLE clients")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (suite *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(suite.client)
	err := suite.accountDB.Create(account)
	suite.Nil(err)
}

func (suite *AccountDBTestSuite) TestFindByID() {
	suite.db.Exec("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		suite.client.ID,
		suite.client.Name,
		suite.client.Email,
		suite.client.CreatedAt,
		suite.client.UpdatedAt)

	account := entity.NewAccount(suite.client)
	err := suite.accountDB.Create(account)
	suite.Nil(err)

	accountDB, err := suite.accountDB.FindByID(account.ID)
	suite.Nil(err)
	suite.NotNil(accountDB)
	suite.Equal(account.ID, accountDB.ID)
	// suite.Equal(account.ClientID, accountDB.ClientID)
	suite.Equal(account.Client.ID, accountDB.Client.ID)
	suite.Equal(account.Client.Name, accountDB.Client.Name)
	suite.Equal(account.Client.Email, accountDB.Client.Email)
	suite.Equal(account.Balance, accountDB.Balance)
}
