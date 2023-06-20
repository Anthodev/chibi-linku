package chibi_linku

import (
	"net/http"
)

func getRootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendResponse(w, "Hello, World!")

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(""))
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
		}
	}
}

func getEncodeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := parseRequest(r.Body, w)
		encodedUrl := Base62Encode(data.Link)

		sendResponse(w, encodedUrl)
	}
}
