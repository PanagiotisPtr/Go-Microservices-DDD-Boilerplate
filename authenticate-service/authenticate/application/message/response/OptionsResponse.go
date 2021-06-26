package response

import (
	"context"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func GetOptionsResponseEncoder(accessibleEndpoints ...string) httptransport.EncodeResponseFunc {
	endpoints := "OPTIONS"
	for _, endpoint := range accessibleEndpoints {
		endpoints += fmt.Sprintf(", %s", endpoint)
	}

	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		w.Header().Set("Allow", endpoints)

		return nil
	}
}
