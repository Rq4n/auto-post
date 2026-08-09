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

	req := httptest.NewRequest("GET", "api/posts", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("handler should not be called")
	})
	AuthMiddleware(handler).ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rec.Code)
	}
}
