package chibi_linku

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidHttpMethod(t *testing.T) {
	nr := httptest.NewRequest(http.MethodPost, "/ping", nil)

	wr := httptest.NewRecorder()

	pingHandler().ServeHTTP(wr, nr)

	if status := wr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}
