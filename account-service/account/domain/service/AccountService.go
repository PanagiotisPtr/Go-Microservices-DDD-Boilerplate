package service

import (
	"account-service/account/domain/dto"
	"account-service/account/domain/irepository"
	"context"

	"github.com/go-kit/kit/log"
)

type AccountService struct {
	repo   irepository.IAccountRepository
	logger log.Logger
}

func NewAccountService(repository irepository.IAccountRepository, logger log.Logger) AccountService {
	return AccountService{
		repo:   repository,
		logger: logger,
	}
}

func (s *AccountService) CreateAccount(ctx context.Context, email string, password string) (string, error) {
	accountUuid, err := s.repo.CreateAccount(ctx, email, password)
	if err != nil {
		return "", err
	}

	return accountUuid, nil
}

func (s *AccountService) GetAccount(ctx context.Context, uuid string) (dto.AccountDto, error) {
	account, err := s.repo.GetAccount(ctx, uuid)
	if err != nil {
		return dto.AccountDto{}, err
	}

	return account, nil
}
