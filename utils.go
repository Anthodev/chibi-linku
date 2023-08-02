package chibi_linku

import (
	"bytes"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"io"
	"net/http"
)

type jer struct {
	Link string `json:"link"`
}

func sendResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "application/json")

	jsonData := buildResponse(response)

	w.WriteHeader(http.StatusOK)
	_, err := w.Write(jsonData)

	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func parseRequest(r io.ReadCloser, w http.ResponseWriter) Url {
	body, err := io.ReadAll(r)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}

	var ur Url
	err = json.Unmarshal(body, &ur)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
	}

	if !govalidator.IsURL(ur.Link) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	return ur
}

func buildResponse(encodedUrl string) []byte {
	jsonValue := jer{
		Link: encodedUrl,
	}

	jsonData, err := json.Marshal(jsonValue)
	jsonData = bytes.TrimPrefix(jsonData, []byte("\xef\xbb\xbf"))

	if err != nil {
		return nil
	}

	return jsonData
}
