package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthRoute(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body["status"] != "ok" {
		t.Fatalf("expected status payload %q, got %#v", "ok", body["status"])
	}
}

func TestCreateUserValidationFailure(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{"name":"Ada"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if _, ok := body["error"]; !ok {
		t.Fatalf("expected error field in response, got %#v", body)
	}
}

func TestSearchDefaultPageBehavior(t *testing.T) {
	router := setupRouter()

	testCases := []struct {
		name string
		url  string
	}{
		{name: "missing page", url: "/search?q=widgets"},
		{name: "invalid page", url: "/search?q=widgets&page=oops"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.url, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
			}

			var body map[string]any
			if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if body["page"] != float64(1) {
				t.Fatalf("expected default page 1, got %#v", body["page"])
			}
		})
	}
}

func TestCreateItemCalculatesTotalAndTax(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest(
		http.MethodPost,
		"/items",
		bytes.NewBufferString(`{"title":"Widget","quantity":3,"price":19.99}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}

	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body["total"] != 59.97 {
		t.Fatalf("expected total 59.97, got %#v", body["total"])
	}

	if body["tax"] != float64(6) {
		t.Fatalf("expected tax 6, got %#v", body["tax"])
	}
}