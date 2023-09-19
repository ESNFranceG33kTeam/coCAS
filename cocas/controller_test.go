package cocas

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	// /login
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	cocasmwd.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	// /logout
	req = httptest.NewRequest(http.MethodGet, "/logout", nil)
	w = httptest.NewRecorder()
	cocasmwd.ServeHTTP(w, req)

	res = w.Result()
	defer res.Body.Close()
	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	// /healthcheck
	req = httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	w = httptest.NewRecorder()
	cocasmwd.ServeHTTP(w, req)

	res = w.Result()
	defer res.Body.Close()
	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	// /conf
	req = httptest.NewRequest(http.MethodGet, "/conf", nil)
	w = httptest.NewRecorder()
	cocasmwd.ServeHTTP(w, req)

	res = w.Result()
	defer res.Body.Close()
	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	// /reload
	req = httptest.NewRequest(http.MethodGet, "/conf", nil)
	w = httptest.NewRecorder()
	cocasmwd.ServeHTTP(w, req)

	res = w.Result()
	defer res.Body.Close()
	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}
