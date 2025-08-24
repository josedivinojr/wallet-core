package createaccount

import (
	"errors"
	"josedivinojr/wallet/internal/entity"
	"josedivinojr/wallet/internal/gateway"
)

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
		ClientGateway:  c,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, errors.New("client not found")
	}

	account := entity.NewAccount(client)
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
