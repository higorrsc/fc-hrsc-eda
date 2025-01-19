package createtransaction

import (
	"github.com/higorrsc/fc-hrsc-microservices/internal/entity"
	"github.com/higorrsc/fc-hrsc-microservices/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountFromID string
	AccountToID   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	transactionGateway gateway.TransactionGateway
	accountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(t gateway.TransactionGateway, a gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		transactionGateway: t,
		accountGateway:     a,
	}
}

func (c *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := c.accountGateway.FindByID(input.AccountFromID)
	if err != nil {
		return nil, err
	}

	accountTo, err := c.accountGateway.FindByID(input.AccountToID)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = c.transactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}, nil
}
