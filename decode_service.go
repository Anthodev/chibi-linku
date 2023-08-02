package chibi_linku

import (
	"github.com/anthodev/chibi_linku/database"
	"net/http"
)

func handleDecodeRequest(w http.ResponseWriter, r *http.Request, s string) {
	rds := database.CreateClient(1)
	decodedUrl, err := database.GetUrl(rds, s)

	if err != nil {
		http.Error(w, "Error getting url from database", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, decodedUrl, http.StatusTemporaryRedirect)
}
