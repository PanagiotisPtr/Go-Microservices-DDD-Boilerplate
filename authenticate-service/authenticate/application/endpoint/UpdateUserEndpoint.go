package endpoint

import (
	"authenticate-service/authenticate/application/message/request"
	"authenticate-service/authenticate/application/message/response"
	"authenticate-service/authenticate/domain/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func UpdateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		requestObject := req.(request.UpdateUserRequest)
		errorResponse := response.UpdateUserResponse{Success: false}

		userUuid, err := s.GetUserUuidFromToken(requestObject.RefreshToken)

		if err != nil {
			return errorResponse, err
		}

		err = s.UpdateUser(
			userUuid,
			requestObject.OldPassword,
			requestObject.NewPassword,
		)

		if err != nil {
			return errorResponse, err
		}

		err = s.RevokeRefreshTokensForUser(requestObject.RefreshToken)

		if err != nil {
			return errorResponse, err
		}

		return response.UpdateUserResponse{Success: true}, err
	}
}
