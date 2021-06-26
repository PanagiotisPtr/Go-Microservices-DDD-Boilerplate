package middleware

import (
	"context"
	"net/http"
	"strings"
)

func AddRequestOriginToContext(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, "Origin", r.Header.Get("Origin"))
}

func AddResponseCorsOptions(ctx context.Context, w http.ResponseWriter) context.Context {
	origin, ok := ctx.Value("Origin").(string)

	if !ok {
		return ctx
	}

	if strings.Contains(origin, "localhost") {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}

	return ctx
}

func AddControlAllowCredentialsHeader(ctx context.Context, w http.ResponseWriter) context.Context {
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	return ctx
}
