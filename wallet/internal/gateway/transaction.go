package gateway

import "josedivinojr/wallet/internal/entity"

type TransactionGateway interface {
	Save(transaction *entity.Transaction) error
}
