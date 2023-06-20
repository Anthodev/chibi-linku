package chibi_linku

import (
	"log"
	"net/http"
)

type Url struct {
	Link       string `json:"link"`
	Expiration int    `json:"expiration"`
}

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRootHandler())
	go mux.HandleFunc("/encode", getEncodeHandler())
	go mux.HandleFunc("/decode/{code}", decodeHandler())

	log.Println("Listening on port 80")

	err := http.ListenAndServe(":8001", mux)

	if err != nil {
		return
	}
}
