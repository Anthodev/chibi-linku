package chibi_linku

import "net/http"

func handleDecodeRequest(w http.ResponseWriter, r *http.Request, s string) {
	decodedUrl := Base62Decode(s)

	http.Redirect(w, r, decodedUrl, http.StatusTemporaryRedirect)
}
