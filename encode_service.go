package chibi_linku

import (
	"github.com/anthodev/chibi_linku/database"
	"github.com/anthodev/chibi_linku/helpers"
	"net/http"
)

func handleEncodeRequest(w http.ResponseWriter, s string, e int) {
	encodedUrl := helpers.Base62Encode(s)

	r := database.CreateClient(1)

	err := database.SaveUrl(r, encodedUrl, s, e)
	if err != nil {
		return
	}

	sendResponse(w, encodedUrl)
}
