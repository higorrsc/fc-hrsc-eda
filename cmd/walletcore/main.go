package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/higorrsc/fc-hrsc-eda/internal/database"
	"github.com/higorrsc/fc-hrsc-eda/internal/event"
	"github.com/higorrsc/fc-hrsc-eda/internal/usecase/create_account"
	"github.com/higorrsc/fc-hrsc-eda/internal/usecase/create_client"
	"github.com/higorrsc/fc-hrsc-eda/internal/usecase/create_transaction"
	"github.com/higorrsc/fc-hrsc-eda/internal/web"
	"github.com/higorrsc/fc-hrsc-eda/internal/web/webserver"
	"github.com/higorrsc/fc-hrsc-eda/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/client", clientHandler.CreateClient)
	webserver.AddHandler("/account", accountHandler.CreateAccount)
	webserver.AddHandler("/transaction", transactionHandler.CreateTransaction)

	webserver.Start()
}
