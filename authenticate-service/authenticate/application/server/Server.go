package server

import (
	applicationEndpoint "authenticate-service/authenticate/application/endpoint"
	"authenticate-service/authenticate/application/message/request"
	"authenticate-service/authenticate/application/message/response"
	"authenticate-service/authenticate/application/middleware"
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints map[string]endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Use(addJsonHeaders)

	r.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		endpoints["RegisterUserEndpoint"],
		request.DecodeRegisterUserRequest,
		response.EncodeRegisterUserResponse,
		httptransport.ServerBefore(middleware.AddRequestOriginToContext),
		httptransport.ServerAfter(middleware.AddResponseCorsOptions),
	))

	r.Methods("OPTIONS").Path("/register").Handler(httptransport.NewServer(
		applicationEndpoint.OptionsEndpoint(),
		request.DecodeOptionsRequest,
		response.GetOptionsResponseEncoder("POST"),
		httptransport.ServerBefore(middleware.AddRequestOriginToContext),
		httptransport.ServerAfter(middleware.AddResponseCorsOptions),
	))

	r.Methods("POST").Path("/authenticate").Handler(httptransport.NewServer(
		endpoints["AuthenticateEndpoint"],
		request.DecodeAuthenticateUserRequest,
		response.EncodeAuthenticateUserResponse,
		httptransport.ServerBefore(middleware.AddRequestOriginToContext),
		httptransport.ServerAfter(middleware.AddResponseCorsOptions),
		httptransport.ServerAfter(middleware.AddControlAllowCredentialsHeader),
	))

	r.Methods("OPTIONS").Path("/authenticate").Handler(httptransport.NewServer(
		applicationEndpoint.OptionsEndpoint(),
		request.DecodeOptionsRequest,
		response.GetOptionsResponseEncoder("POST"),
		httptransport.ServerBefore(middleware.AddRequestOriginToContext),
		httptransport.ServerAfter(middleware.AddResponseCorsOptions),
	))

	r.Methods("GET").Path("/get_jwt").Handler(httptransport.NewServer(
		endpoints["GetJwtEndpoint"],
		request.DecodeGetJwtRequest,
		response.EncodeGetJwtResponse,
		httptransport.ServerBefore(middleware.AddRequestOriginToContext),
		httptransport.ServerAfter(middleware.AddResponseCorsOptions),
		httptransport.ServerAfter(middleware.AddControlAllowCredentialsHeader),
	))

	r.Methods("OPTIONS").Path("/get_jwt").Handler(httptransport.NewServer(
		applicationEndpoint.OptionsEndpoint(),
		request.DecodeOptionsRequest,
		response.GetOptionsResponseEncoder("GET"),
		httptransport.ServerBefore(middleware.AddRequestOriginToContext),
		httptransport.ServerAfter(middleware.AddResponseCorsOptions),
	))

	r.Methods("POST").Path("/logout").Handler(httptransport.NewServer(
		endpoints["LogoutUserEndpoint"],
		request.DecodeLogoutUserRequest,
		response.EncodeLogoutUserResponse,
		httptransport.ServerBefore(middleware.AddRequestOriginToContext),
		httptransport.ServerAfter(middleware.AddResponseCorsOptions),
		httptransport.ServerAfter(middleware.AddControlAllowCredentialsHeader),
	))

	r.Methods("OPTIONS").Path("/logout").Handler(httptransport.NewServer(
		applicationEndpoint.OptionsEndpoint(),
		request.DecodeOptionsRequest,
		response.GetOptionsResponseEncoder("POST"),
		httptransport.ServerBefore(middleware.AddRequestOriginToContext),
		httptransport.ServerAfter(middleware.AddResponseCorsOptions),
	))

	r.Methods("PUT").Path("/change_password").Handler(httptransport.NewServer(
		endpoints["UpdateUserEndpoint"],
		request.DecodeUpdateUserRequest,
		response.EncodeUpdateUserResponse,
		httptransport.ServerBefore(middleware.AddRequestOriginToContext),
		httptransport.ServerAfter(middleware.AddResponseCorsOptions),
		httptransport.ServerAfter(middleware.AddControlAllowCredentialsHeader),
	))

	r.Methods("OPTIONS").Path("/change_password").Handler(httptransport.NewServer(
		applicationEndpoint.OptionsEndpoint(),
		request.DecodeOptionsRequest,
		response.GetOptionsResponseEncoder("PUT"),
		httptransport.ServerBefore(middleware.AddRequestOriginToContext),
		httptransport.ServerAfter(middleware.AddResponseCorsOptions),
	))

	return r
}

func addJsonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
