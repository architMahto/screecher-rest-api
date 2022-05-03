package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/customers", nil)
	res := httptest.NewRecorder()

	HandleNotFound(res, req)

	if res.Code != http.StatusNotFound {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusNotFound)
	}
}
