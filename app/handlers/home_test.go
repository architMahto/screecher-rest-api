package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/architMahto/screecher-rest-api/app/handlers"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	handlers.Home(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusOK)
	}
}
