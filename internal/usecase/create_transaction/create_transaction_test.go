package createtransaction

import (
	"testing"

	"github.com/higorrsc/fc-hrsc-microservices/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "L4uQK@example.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000.00)

	client2, _ := entity.NewClient("Jane Doe", "onpen@huitiho.lv")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000.00)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindByID", account1.ID).Return(account1, nil)
	mockAccount.On("FindByID", account2.ID).Return(account2, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountFromID: account1.ID,
		AccountToID:   account2.ID,
		Amount:        100.00,
	}

	uc := NewCreateTransactionUseCase(mockTransaction, mockAccount)
	outputDto, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, outputDto.ID)
	mockAccount.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindByID", 2)
	mockTransaction.AssertExpectations(t)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
