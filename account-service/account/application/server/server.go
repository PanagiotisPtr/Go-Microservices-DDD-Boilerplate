package server

import (
	"account-service/account/application/message/request"
	"account-service/account/application/message/response"
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints map[string]endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Use(addJsonHeaders)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints["CreateAccountEndpoint"],
		request.DecodeCreateAccountRequest,
		response.EncodeCreateAccountResponse,
	))

	r.Methods("GET").Path("/user/{uuid}").Handler(httptransport.NewServer(
		endpoints["GetAccountEndpoint"],
		request.DecodeGetAccountRequest,
		response.EncodeGetAccountResponse,
	))

	return r
}

func addJsonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
