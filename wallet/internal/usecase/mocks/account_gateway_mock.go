package mocks

import (
	"josedivinojr/wallet/internal/entity"

	"github.com/stretchr/testify/mock"
)

type AccountGateway struct {
	mock.Mock
}

func (m *AccountGateway) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGateway) GetById(id string) (*entity.Account, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Account), args.Error(1)
}
