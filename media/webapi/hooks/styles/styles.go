package styles

import (
	"net/http"
)

type AddressMap = map[string]string

const baseboard = "/styles/baseboard.css"
const webClient = "/styles/web_client.css"
const webClientLogin = "/styles/web_client_login.css"
const address = "/usr/local/web_clients/web_client/dist"
const addressLogin = "/usr/local/web_clients/web_client_login/dist"

var addresses = AddressMap{
	baseboard: address + baseboard,
	webClient: address + webClient,
	webClientLogin: address + webClientLogin,
}

// Styles -
func Styles(w http.ResponseWriter, r *http.Request) {
	filePath, found := addresses[r.URL.Path]
	if !found {
		return
	}
	
	http.ServeFile(w, r, filePath)
}