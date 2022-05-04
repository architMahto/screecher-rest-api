package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/architMahto/screecher-rest-api/app/handlers"
)

func TestHandleNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/customers", nil)
	res := httptest.NewRecorder()

	handlers.HandleNotFound(res, req)

	if res.Code != http.StatusNotFound {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusNotFound)
	}
}
