package chibi_linku

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type er struct {
	Link string `json:"link"`
}

func TestWebServerRunning(t *testing.T) {
	nr := httptest.NewRequest("GET", "/", nil)

	wr := httptest.NewRecorder()

	getRootHandler().ServeHTTP(wr, nr)

	if status := wr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `"Hello, World!"`

	if wr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			wr.Body.String(), expected)
	}
}

func TestEncoreUrl(t *testing.T) {
	url := Url{
		Link:       "https://google.com",
		Expiration: 0,
	}

	bodyValue, _ := json.Marshal(url)

	nr := httptest.NewRequest("POST", "/encode", bytes.NewBuffer(bodyValue))
	nr.Header.Set("Content-Type", "application/json")

	wr := httptest.NewRecorder()

	getEncodeHandler().ServeHTTP(wr, nr)

	if status := wr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	responseValue := er{}
	err := json.Unmarshal(wr.Body.Bytes(), &responseValue)

	if err != nil {
		t.Errorf("cannot unmarshal response: %v", err)
	}

	expected := er{
		Link: "2my0rvHAvGxZOpU2jvJcLJw1",
	}

	if responseValue != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			wr.Body.String(), expected)
	}
}

func TestDecodeAndRedirect(t *testing.T) {
	nr := httptest.NewRequest("GET", "/decode/2my0rvHAvGxZOpU2jvJcLJw1", nil)
	wr := httptest.NewRecorder()

	decodeHandler().ServeHTTP(wr, nr)

	if status := wr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusTemporaryRedirect)
	}

	expected := "https://google.com"

	if wr.Header().Get("Location") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", wr.Body.String(), expected)
	}
}
