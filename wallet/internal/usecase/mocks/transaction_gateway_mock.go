package mocks

import (
	"josedivinojr/wallet/internal/entity"

	"github.com/stretchr/testify/mock"
)

type TransactionGateway struct {
	mock.Mock
}

func (m *TransactionGateway) Save(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}
