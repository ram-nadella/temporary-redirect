package temporaryredirect

import (
	"net/http"
)

func init() {
	redirectHandler := http.RedirectHandler("https://307.temporaryredirect.com", http.StatusTemporaryRedirect)
	http.Handle("/", redirectHandler)
}
