package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

func TestAuthMiddleware_NoSession(t *testing.T) {
	// new cookie store because the auth middleware depends on a cookie store
	// so we create a mock one
	gothic.Store = sessions.NewCookieStore([]byte("test-secret"))

	req := httptest.NewRequest("GET", "/api/post", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("handler should not be called")
	})
	AuthMiddleware(handler).ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rec.Code)
	}
}

func TestAuthMiddleware_Authenticated(t *testing.T) {
	gothic.Store = sessions.NewCookieStore([]byte("test-secret"))

	req := httptest.NewRequest("GET", "/api/post", nil)
	rec := httptest.NewRecorder()

	// get session value and save like cmd/callback.go
	session, _ := gothic.Store.Get(req, "session")
	session.Values["user_id"] = "123"
	session.Save(req, rec)

	req.Header.Set("Cookie", rec.Header().Get("Set-cookie"))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get userID context
		// if received userID != got

		userID := r.Context().Value("userID")

		if userID != "123" {
			t.Errorf("expected userID 123 got %v", userID)
		}

		w.WriteHeader(http.StatusOK)
	})
	AuthMiddleware(handler).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}

func TestAuthMiddleware_NoUserID(t *testing.T) {
	gothic.Store = sessions.NewCookieStore([]byte("test-secret"))

	req := httptest.NewRequest("GET", "/api/post", nil)
	rec := httptest.NewRecorder()

	session, _ := gothic.Store.Get(req, "session")
	session.Save(req, rec)

	req.Header.Set("Cookie", rec.Header().Get("Set-cookie"))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("handler should not be called")
	})
	AuthMiddleware(handler).ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected 401 got %d", rec.Code)
	}
}
