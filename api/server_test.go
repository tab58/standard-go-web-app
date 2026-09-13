package api_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"standard-go-web-app/api"
	"standard-go-web-app/internal/app"
)

// The library's Start binds internally without exposing the address, so tests
// use a fixed port on localhost.
const testAddr = "127.0.0.1:18765"

func startTestServer(t *testing.T) string {
	t.Helper()
	srv := api.NewServer(api.ServerConfig{
		ServiceName:    "standard-go-web-app-test",
		ServiceVersion: "test",
		JWTSecret:      "test-secret-that-is-long-enough",
	}, app.New())
	errCh, err := srv.Start(testAddr)
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}
	t.Cleanup(func() { srv.Shutdown(t.Context()) })
	go func() {
		if err := <-errCh; err != nil {
			t.Errorf("server error: %v", err)
		}
	}()
	return "http://" + testAddr
}

func TestHealth(t *testing.T) {
	base := startTestServer(t)
	resp, err := http.Get(base + "/health")
	if err != nil {
		t.Fatalf("GET /health error = %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("GET /health status = %d, want 200", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), `"ok"`) {
		t.Errorf("GET /health body = %q", body)
	}
}
