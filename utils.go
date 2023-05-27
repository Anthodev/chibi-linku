package chibi_linku

import (
	"encoding/json"
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)

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

	return ur
}
