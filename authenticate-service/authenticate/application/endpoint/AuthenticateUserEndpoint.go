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
		errorResponse := response.AuthenticateUserResponse{
			Success:      false,
			Token:        "",
			RefreshToken: "",
		}

		requestObject := req.(request.AuthenticateUserRequest)
		refreshToken, err := s.AuthenticateUser(requestObject.Email, requestObject.Password)

		if err != nil {
			return errorResponse, err
		}

		token, err := s.GetJWT(refreshToken)

		if err != nil {
			return errorResponse, err
		}

		return response.AuthenticateUserResponse{
			Success:      true,
			Token:        token,
			RefreshToken: refreshToken,
		}, nil
	}
}
