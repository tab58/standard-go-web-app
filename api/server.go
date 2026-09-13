// Package api wires the standard-go-web-app application onto the
// huma-http-server library: server creation, route registration, and
// request/response types.
package api

import (
	"net/http"

	server "github.com/tab58/huma-http-server"
	"github.com/tab58/huma-http-server/router"

	"standard-go-web-app/internal/app"
)

// ServerConfig carries the settings NewServer needs.
type ServerConfig struct {
	ServiceName    string
	ServiceVersion string
	JWTSecret      string
}

// NewServer wires the application: auth middleware (raw claims, no typed
// user object) and the raw /health route. Domain routes are registered here
// as the control plane gains resources.
func NewServer(cfg ServerConfig, application *app.Application) *server.Server[router.MapAuthInfo] {
	srv := server.New(server.ServerConfig{
		ServiceName:      cfg.ServiceName,
		ServiceVersion:   cfg.ServiceVersion,
		JWTSigningSecret: cfg.JWTSecret,
	}, router.MapAuthInfoBuilder)

	_ = application // domain routes will consume the injected Application

	// raw route: bypasses middleware, not in OpenAPI
	srv.Handle("/healthz", healthHandler())

	return srv
}

func healthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
}
