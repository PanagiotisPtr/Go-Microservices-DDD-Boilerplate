package endpoint

import (
	"account-service/account/application/message/request"
	"account-service/account/application/message/response"
	"account-service/account/domain/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func CreateAccountEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		requestObject := req.(request.CreateAccountRequest)
		uuid, err := s.CreateAccount(ctx, requestObject.Email, requestObject.Password)
		if err != nil {
			return response.CreateAccountResponse{}, err
		}

		return response.CreateAccountResponse{Uuid: uuid}, err
	}
}
