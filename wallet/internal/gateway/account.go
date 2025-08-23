package gateway

import "josedivinojr/wallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	GetById(id string) (*entity.Account, error)
}
