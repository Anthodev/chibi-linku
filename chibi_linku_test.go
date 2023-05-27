package chibi_linku

import (
	"net/http/httptest"
	"testing"
)

func TestWebServerRunning(t *testing.T) {
	nr := httptest.NewRequest("GET", "/", nil)

	wr := httptest.NewRecorder()

	getRootHandler()(wr, nr)

	if wr.Code != 200 {
		t.Errorf("Expected status code 200, got %d", wr.Code)
	}

	if wr.Body.String() != "Hello, World!" {
		t.Errorf("Expected body 'Hello, World!', got %s", wr.Body.String())
	}
}
