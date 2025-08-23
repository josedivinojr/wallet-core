package createclient

import (
	"errors"
	"testing"

	"josedivinojr/wallet/internal/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute_Success(t *testing.T) {
	clientGateway := &ClientGatewayMock{}
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

func TestCreateClientUseCase_Execute_Error(t *testing.T) {
	clientGateway := &ClientGatewayMock{}
	clientGateway.On("Save", mock.Anything).Return(errors.New("error saving client"))

	uc := NewCreateClientUseCase(clientGateway)

	input := CreateClientInputDTO{
		Name:  "John",
		Email: "j@j.com",
	}

	output, err := uc.Execute(input)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.Equal(t, "error saving client", err.Error())

	clientGateway.AssertExpectations(t)
	clientGateway.AssertNumberOfCalls(t, "Save", 1)
}
