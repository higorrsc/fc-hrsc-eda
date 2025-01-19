package database

import (
	"database/sql"
	"testing"

	"github.com/higorrsc/fc-hrsc-microservices/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (suite *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255) PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), created_at DATETIME, updated_at DATETIME)")
	suite.clientDB = NewClientDb(db)
}

func (suite *ClientDBTestSuite) TearDownSuite() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (suite *ClientDBTestSuite) TestSave() {
	client, _ := entity.NewClient("John Doe", "L4uQK@example.com")
	err := suite.clientDB.Save(client)
	suite.Nil(err)
}

func (suite *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("John Doe", "L4uQK@example.com")
	suite.clientDB.Save(client)

	clientDB, err := suite.clientDB.Get(client.ID)
	suite.Nil(err)
	suite.NotNil(clientDB)
	suite.Equal(client.ID, clientDB.ID)
	suite.Equal(client.Name, clientDB.Name)
	suite.Equal(client.Email, clientDB.Email)
}
