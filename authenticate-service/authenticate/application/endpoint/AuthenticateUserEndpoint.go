package endpoint

import (
	"authenticate-service/authenticate/application/message/request"
	"authenticate-service/authenticate/application/message/response"
	"authenticate-service/authenticate/domain/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func AuthenticateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		requestObject := req.(request.AuthenticateUserRequest)
		refreshToken, err := s.AuthenticateUser(requestObject.Email, requestObject.Password)

		return response.AuthenticateUserResponse{Success: (err == nil), Token: refreshToken}, err
	}
}
