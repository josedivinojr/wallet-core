package mocks

import (
	"josedivinojr/wallet/internal/entity"

	"github.com/stretchr/testify/mock"
)

type ClientGateway struct {
	mock.Mock
}

func (m *ClientGateway) Get(id string) (*entity.Client, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGateway) Save(client *entity.Client) error {
	args := m.Called(client)

	return args.Error(0)
}
