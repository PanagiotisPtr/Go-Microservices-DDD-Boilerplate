package endpoint

import (
	"account-service/account/application/message/request"
	"account-service/account/application/message/response"
	"account-service/account/domain/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func GetAccountEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		requestObject := req.(request.GetAccountRequest)
		account, err := s.GetAccount(ctx, requestObject.Uuid)
		if err != nil {
			return response.GetAccountResponse{}, err
		}

		return response.GetAccountResponse{
			Uuid:     account.Uuid,
			Email:    account.Email,
			Password: account.Password,
		}, nil
	}
}
