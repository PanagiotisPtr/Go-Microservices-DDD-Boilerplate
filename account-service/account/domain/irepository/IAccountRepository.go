package irepository

import (
	"account-service/account/domain/dto"
	"context"
)

type IAccountRepository interface {
	CreateAccount(ctx context.Context, email string, password string) (string, error)
	GetAccount(ctx context.Context, uuid string) (dto.AccountDto, error)
}
