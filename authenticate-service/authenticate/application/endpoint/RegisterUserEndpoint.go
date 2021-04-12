package endpoint

import (
	"authenticate-service/authenticate/application/message/request"
	"authenticate-service/authenticate/application/message/response"
	"authenticate-service/authenticate/domain/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func RegisterUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		requestObject := req.(request.RegisterUserRequest)
		err := s.RegisterUser(requestObject.Email, requestObject.Password)

		return response.RegisterUserResponse{Success: (err == nil)}, err
	}
}
