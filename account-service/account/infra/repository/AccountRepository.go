package repository

import (
	"account-service/account/domain/dto"
	"account-service/account/domain/irepository"
	"account-service/account/infra/entity"
	"context"

	"github.com/go-kit/kit/log"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewAccountRepository(db *gorm.DB, logger log.Logger) irepository.IAccountRepository {
	return &AccountRepository{
		db:     db,
		logger: log.With(logger, "AccountRepository", "PostgreSQL"),
	}
}

func (repo *AccountRepository) CreateAccount(ctx context.Context, email string, password string) (string, error) {
	accountUuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	account := entity.Account{
		Uuid:     accountUuid.String(),
		Email:    email,
		Password: password,
	}

	result := repo.db.Scopes(entity.AccountTable(account)).Create(&account)
	if result.Error != nil {
		return "", result.Error
	}

	return account.Uuid, nil
}

func (repo *AccountRepository) GetAccount(ctx context.Context, uuid string) (dto.AccountDto, error) {
	var account entity.Account
	result := repo.db.Scopes(entity.AccountTable(account)).First(&account, "uuid = ?", uuid)

	if result.Error != nil {
		return dto.AccountDto{}, result.Error
	}

	return dto.AccountDto{
		Uuid:     account.Uuid,
		Email:    account.Email,
		Password: account.Password,
	}, nil
}
