package createclient

import (
	"testing"

	"josedivinojr/wallet/internal/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	clientGateway := &mocks.ClientGateway{}
	clientGateway.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(clientGateway)

	input := CreateClientInputDTO{
		Name:  "John",
		Email: "j@j.com",
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Email, output.Email)

	clientGateway.AssertExpectations(t)
	clientGateway.AssertNumberOfCalls(t, "Save", 1)
}
