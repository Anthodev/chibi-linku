package chibi_linku

import (
	"bytes"
	"encoding/json"
	"github.com/anthodev/chibi_linku/database"
	"net/http"
	"net/http/httptest"
	"testing"
)

type er struct {
	Link string `json:"link"`
}

func TestWebServerRunning(t *testing.T) {
	nr := httptest.NewRequest(http.MethodGet, "/ping", nil)

	wr := httptest.NewRecorder()

	pingHandler().ServeHTTP(wr, nr)

	if status := wr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `"Pong!"`

	if wr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			wr.Body.String(), expected)
	}
}

func TestEncodeUrl(t *testing.T) {
	rdb := database.CreateClient(1)
	err := database.FlushAll(rdb)

	url := Url{
		Link:       "https://google.com",
		Expiration: 0,
	}

	bodyValue, _ := json.Marshal(url)
	bodyValue = bytes.TrimPrefix(bodyValue, []byte("\xef\xbb\xbf"))

	nr := httptest.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer(bodyValue))
	nr.Header.Set("Content-Type", "application/json")

	wr := httptest.NewRecorder()

	encodeHandler().ServeHTTP(wr, nr)

	if status := wr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	responseValue := er{}
	err = json.Unmarshal(wr.Body.Bytes(), &responseValue)

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

func TestEncodeUrlWithExpiration(t *testing.T) {
	rdb := database.CreateClient(1)
	err := database.FlushAll(rdb)

	url := Url{
		Link:       "https://google.com",
		Expiration: 3600,
	}

	bodyValue, _ := json.Marshal(url)
	bodyValue = bytes.TrimPrefix(bodyValue, []byte("\xef\xbb\xbf"))

	nr := httptest.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer(bodyValue))
	nr.Header.Set("Content-Type", "application/json")

	wr := httptest.NewRecorder()

	encodeHandler().ServeHTTP(wr, nr)

	if status := wr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	responseValue := er{}
	err = json.Unmarshal(wr.Body.Bytes(), &responseValue)

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
	rdb := database.CreateClient(1)
	err := database.FlushAll(rdb)
	err = database.SaveUrl(rdb, "2my0rvHAvGxZOpU2jvJcLJw1", "https://google.com", 0)

	if err != nil {
		return
	}

	nr := httptest.NewRequest(http.MethodGet, "/decode/2my0rvHAvGxZOpU2jvJcLJw1", nil)
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

func TestDecodeAndRedirectWithExpiration(t *testing.T) {
	rdb := database.CreateClient(1)
	err := database.FlushAll(rdb)
	err = database.SaveUrl(rdb, "2my0rvHAvGxZOpU2jvJcLJw1", "https://google.com", 3600)

	if err != nil {
		return
	}

	nr := httptest.NewRequest(http.MethodGet, "/decode/2my0rvHAvGxZOpU2jvJcLJw1", nil)
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

func TestDecodeAndRedirectWithExpiredUrl(t *testing.T) {
	rdb := database.CreateClient(1)
	err := database.FlushAll(rdb)
	err = database.SaveUrl(rdb, "2my0rvHAvGxZOpU2jvJcLJw1", "https://google.com", -3600)

	if err != nil {
		return
	}

	nr := httptest.NewRequest(http.MethodGet, "/decode/2my0rvHAvGxZOpU2jvJcLJw1", nil)
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
