package server

import (
	"authenticate-service/authenticate/application/message/request"
	"authenticate-service/authenticate/application/message/response"
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
		endpoints["RegisterUserEndpoint"],
		request.DecodeRegisterUserRequest,
		response.EncodeRegisterUserResponse,
	))

	r.Methods("POST").Path("/authenticate").Handler(httptransport.NewServer(
		endpoints["AuthenticateEndpoint"],
		request.DecodeAuthenticateUserRequest,
		response.EncodeAuthenticateUserResponse,
	))

	r.Methods("GET").Path("/get_jwt").Handler(httptransport.NewServer(
		endpoints["GetJwtEndpoint"],
		request.DecodeGetJwtRequest,
		response.EncodeGetJwtResponse,
	))

	r.Methods("POST").Path("/logout").Handler(httptransport.NewServer(
		endpoints["LogoutUserEndpoint"],
		request.DecodeLogoutUserRequest,
		response.EncodeLogoutUserResponse,
	))

	r.Methods("PUT").Path("/change_password").Handler(httptransport.NewServer(
		endpoints["UpdateUserEndpoint"],
		request.DecodeUpdateUserRequest,
		response.EncodeUpdateUserResponse,
	))

	return r
}

func addJsonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
