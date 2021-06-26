package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func OptionsEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	}
}
