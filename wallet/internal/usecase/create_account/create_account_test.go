package createaccount

import (
	"testing"

	"josedivinojr/wallet/internal/entity"
	"josedivinojr/wallet/internal/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John", "j@j.com")

	clientGateway := &mocks.ClientGateway{}
	clientGateway.On("Get", client.ID).Return(client, nil)

	accountGateway := &mocks.AccountGateway{}
	accountGateway.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountGateway, clientGateway)

	input := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	clientGateway.AssertExpectations(t)
	accountGateway.AssertExpectations(t)
	clientGateway.AssertNumberOfCalls(t, "Get", 1)
	accountGateway.AssertNumberOfCalls(t, "Save", 1)
}
