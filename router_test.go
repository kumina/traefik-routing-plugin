package traefik_routing_plugin_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kumina/traefik-routing-plugin"
)

func TestRouter(t *testing.T)  {
	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	cfg := traefik_routing_plugin.CreateConfig()
	cfg.Routes["Route-To-Service"] = "Dummy"
	cfg.Routes["Route-To-Service-Test"] = "DummyTest"
	cfg.Routes["Route-To-Service-Lambda"] = "Lambda"

	handler, err := traefik_routing_plugin.New(ctx, next, cfg, "traefik-routing-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "Route-To-Service", "Dummy")
	assertHeader(t, req, "Route-To-Service-Test", "DummyTest")
	assertHeader(t, req, "Route-To-Service-Lambda", "Lambda")

	recorder = httptest.NewRecorder()
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/asd", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "Route-To-Service", "Dummy")
	assertHeader(t, req, "Route-To-Service-Test", "DummyTest")
	assertHeader(t, req, "Route-To-Service-Lambda", "Lambda")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
