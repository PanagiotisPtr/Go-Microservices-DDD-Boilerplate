package endpoint

import (
	"authenticate-service/authenticate/application/message/request"
	"authenticate-service/authenticate/application/message/response"
	"authenticate-service/authenticate/domain/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func LogoutUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		requestObject := req.(request.LogoutUserRequest)
		err := s.RevokeRefreshTokensForUser(requestObject.RefreshToken)

		return response.LogoutUserResponse{Success: (err == nil)}, err
	}
}
