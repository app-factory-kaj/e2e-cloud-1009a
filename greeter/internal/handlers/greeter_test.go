package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"greeter/internal/gen"
	"greeter/internal/handlers"
)

func newRouter() http.Handler {
	r := chi.NewRouter()
	srv := handlers.NewServer()
	return gen.HandlerWithOptions(gen.NewStrictHandler(srv, nil), gen.ChiServerOptions{BaseRouter: r})
}

func TestHealth(t *testing.T) {
	h := newRouter()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var body gen.Health
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body.Status != "ok" {
		t.Fatalf("expected status ok, got %q", body.Status)
	}
}

func TestGreetingWithName(t *testing.T) {
	h := newRouter()
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Ada", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var body gen.Greeting
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body.Message != "Hello, Ada!" {
		t.Fatalf("expected %q, got %q", "Hello, Ada!", body.Message)
	}
}

func TestGreetingDefaultsToWorld(t *testing.T) {
	cases := []string{"/hello", "/hello?name="}
	for _, path := range cases {
		h := newRouter()
		req := httptest.NewRequest(http.MethodGet, path, nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("%s: expected 200, got %d", path, w.Code)
		}
		var body gen.Greeting
		if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
			t.Fatalf("%s: decode: %v", path, err)
		}
		if body.Message != "Hello, World!" {
			t.Fatalf("%s: expected %q, got %q", path, "Hello, World!", body.Message)
		}
	}
}
