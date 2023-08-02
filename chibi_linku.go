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
	go mux.HandleFunc("/encode", encodeHandler())
	go mux.HandleFunc("/decode/{code}", decodeHandler())
	go mux.HandleFunc("/purge", purgeHandler())
	go mux.HandleFunc("/", rootHandler())

	log.Println("Listening on port 80")

	err := http.ListenAndServe(":8001", mux)

	if err != nil {
		return
	}
}
