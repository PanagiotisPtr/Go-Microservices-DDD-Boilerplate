package request

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type GetAccountRequest struct {
	Uuid string `json:"uuid"`
}

func DecodeGetAccountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	uuid, ok := vars["uuid"]
	if !ok {
		return GetAccountRequest{}, errors.New("Missing UUID parameter in path for GetAccount request")
	}

	return GetAccountRequest{
		Uuid: uuid,
	}, nil
}
