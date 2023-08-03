package chibi_linku

import (
	"encoding/json"
	"github.com/anthodev/chibi_linku/database"
	"net/http"
	"strings"
)

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		marshal, err := json.Marshal("Pong!")
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(marshal)
		if err != nil {
			return
		}
	}
}

func encodeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		request := parseRequest(r.Body, w)
		link := request.Link
		expiration := request.Expiration
		handleEncodeRequest(w, link, expiration)
	}
}

func decodeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		vars := strings.Split(r.RequestURI, "/")
		code := vars[len(vars)-1]
		handleDecodeRequest(w, r, code)
	}
}

func purgeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		rds := database.CreateClient(1)
		err := rds.FlushAll(r.Context()).Err()
		if err != nil {
			return
		}

		sendResponse(w, "OK")
	}
}
