package chibi_linku

import "net/http"

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

		_, err := w.Write([]byte(data.Link + " " + string(rune(data.Expiration))))
		if err != nil {
			return
		}
	}
}
