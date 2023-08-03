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
	port := "8001"
	mux := http.NewServeMux()
	go mux.HandleFunc("/encode", encodeHandler())
	go mux.HandleFunc("/decode/{code}", decodeHandler())
	go mux.HandleFunc("/purge", purgeHandler())
	go mux.HandleFunc("/ping", pingHandler())

	log.Println("Listening on port " + port)

	err := http.ListenAndServe(port, mux)

	if err != nil {
		return
	}
}
