package endpoint

import (
	"authenticate-service/authenticate/application/message/request"
	"authenticate-service/authenticate/application/message/response"
	"authenticate-service/authenticate/domain/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func GetJwtEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		errorResponse := response.GetJwtResponse{
			Success: false,
			Token:   "",
		}

		requestObject, ok := req.(request.GetJwtRequest)

		if !ok {
			return errorResponse, nil
		}

		token, err := s.GetJWT(requestObject.RefreshToken)

		if err != nil {
			return errorResponse, nil
		}

		return response.GetJwtResponse{
			Success: true,
			Token:   token,
		}, nil
	}
}
