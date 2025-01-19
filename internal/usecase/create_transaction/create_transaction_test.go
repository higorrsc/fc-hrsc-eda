package create_transaction

import (
	"context"
	"testing"

	"github.com/higorrsc/fc-hrsc-eda/internal/entity"
	"github.com/higorrsc/fc-hrsc-eda/internal/event"
	"github.com/higorrsc/fc-hrsc-eda/internal/usecase/mocks"
	"github.com/higorrsc/fc-hrsc-eda/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "L4uQK@example.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000.00)

	client2, _ := entity.NewClient("Jane Doe", "onpen@huitiho.lv")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000.00)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountFromID: account1.ID,
		AccountToID:   account2.ID,
		Amount:        100.00,
	}

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalance := event.NewBalanceUpdated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransaction, eventBalance)
	outputDto, err := uc.Execute(ctx, inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, outputDto.ID)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
