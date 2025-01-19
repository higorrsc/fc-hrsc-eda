package create_transaction

import (
	"context"

	"github.com/higorrsc/fc-hrsc-eda/internal/entity"
	"github.com/higorrsc/fc-hrsc-eda/internal/gateway"
	"github.com/higorrsc/fc-hrsc-eda/pkg/events"
	"github.com/higorrsc/fc-hrsc-eda/pkg/uow"
)

type CreateTransactionInputDTO struct {
	AccountFromID string  `json:"account_id_from"`
	AccountToID   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	ID            string  `json:"id"`
	AccountFromID string  `json:"account_id_from"`
	AccountToID   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	unitOfWork         uow.UowInterface
	eventDispatcher    events.EventDispatcherInterface
	transactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	uow uow.UowInterface,
	ed events.EventDispatcherInterface,
	tc events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		unitOfWork:         uow,
		eventDispatcher:    ed,
		transactionCreated: tc,
	}
}

func (c *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	output := &CreateTransactionOutputDTO{}

	err := c.unitOfWork.Do(ctx, func(_ *uow.Uow) error {
		accountRepository := c.getAccountRepository(ctx)
		transactionRepository := c.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindByID(input.AccountFromID)
		if err != nil {
			return err
		}

		accountTo, err := accountRepository.FindByID(input.AccountToID)
		if err != nil {
			return err
		}

		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			return err
		}

		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}

		output.ID = transaction.ID
		output.AccountFromID = input.AccountFromID
		output.AccountToID = input.AccountToID
		output.Amount = input.Amount

		return nil
	})

	if err != nil {
		return nil, err
	}

	c.transactionCreated.SetPayload(output)
	c.eventDispatcher.Dispatch(c.transactionCreated)

	return output, nil
}

func (c *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := c.unitOfWork.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (c *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := c.unitOfWork.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
