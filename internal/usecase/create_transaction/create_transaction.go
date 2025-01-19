package create_transaction

import (
	"github.com/higorrsc/fc-hrsc-eda/internal/entity"
	"github.com/higorrsc/fc-hrsc-eda/internal/gateway"
	"github.com/higorrsc/fc-hrsc-eda/pkg/events"
)

type CreateTransactionInputDTO struct {
	AccountFromID string  `json:"account_id_from"`
	AccountToID   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	ID string `json:"id"`
}

type CreateTransactionUseCase struct {
	transactionGateway gateway.TransactionGateway
	accountGateway     gateway.AccountGateway
	eventDispatcher    events.EventDispatcherInterface
	transactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	t gateway.TransactionGateway,
	a gateway.AccountGateway,
	ed events.EventDispatcherInterface,
	tc events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		transactionGateway: t,
		accountGateway:     a,
		eventDispatcher:    ed,
		transactionCreated: tc,
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

	output := &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}

	c.transactionCreated.SetPayload(output)
	c.eventDispatcher.Dispatch(c.transactionCreated)

	return output, nil
}
