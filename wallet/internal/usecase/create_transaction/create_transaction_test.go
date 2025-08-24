package createtransaction

import (
	"testing"

	"josedivinojr/wallet/internal/entity"
	"josedivinojr/wallet/internal/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John", "j@j.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("Marie", "m@m.com")
	account2 := entity.NewAccount(client2)

	accountGateway := &mocks.AccountGateway{}
	accountGateway.On("GetById", account1.ID).Return(account1, nil)
	accountGateway.On("GetById", account2.ID).Return(account2, nil)

	transactionGateway := &mocks.TransactionGateway{}
	transactionGateway.On("Save", mock.Anything).Return(nil)

	uc := NewCreateTransactionUseCase(transactionGateway, accountGateway)

	input := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	accountGateway.AssertExpectations(t)
	transactionGateway.AssertExpectations(t)
	accountGateway.AssertNumberOfCalls(t, "GetById", 2)
	transactionGateway.AssertNumberOfCalls(t, "Save", 1)
}
