// Command app is the standard-go-web-app server entrypoint: it loads
// configuration, constructs the application core, injects it into the api
// server, and runs until interrupted.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"standard-go-web-app/api"
	app "standard-go-web-app/internal/app"

	"standard-go-web-app/cmd/app/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// build the application core and inject it into the api server
	application := app.New()
	srv := api.NewServer(api.ServerConfig{
		ServiceName:    "standard-go-web-app",
		ServiceVersion: "0.1.0",
		JWTSecret:      cfg.JWTSecret,
	}, application)

	errCh, err := srv.Start(cfg.Port)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "standard-go-web-app listening on %s\n", cfg.Port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	select {
	case err := <-errCh:
		fmt.Fprintln(os.Stderr, "server error:", err)
	case <-stop:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}
}
