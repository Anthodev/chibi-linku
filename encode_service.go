package chibi_linku

import "net/http"

func handleEncodeRequest(w http.ResponseWriter, s string) {
	encodedUrl := Base62Encode(s)

	sendResponse(w, encodedUrl)
}
