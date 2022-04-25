package routes_test

import (
	"apiTemplate/internal/api/routes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ping", nil)

	r := routes.SetupRouter()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}
