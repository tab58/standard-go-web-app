// Package app holds the standard-go-web-app application core: the
// Application struct with its dependencies. It has no HTTP knowledge — the
// api package wires this onto the huma-http-server library.
package app

// Application is the application core, dependency-injected into the api
// server. Zero-value is not usable; construct with New.
type Application struct{}

// New creates the Application.
func New() *Application {
	return &Application{}
}
