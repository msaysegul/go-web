package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestFeed(t *testing.T) {
	e := NewEngine()
	req := httptest.NewRequest(echo.GET, "/recipes/feed", nil)
	rec := httptest.NewRecorder()
	e.Echo.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
