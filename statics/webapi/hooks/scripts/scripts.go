package scripts

import (
	"net/http"
)

// AddressMap -
type AddressMap = map[string]string

const webClient = "/scripts/web_client.js"
const webClientLogin = "/scripts/web_client_login.js"
const address = "/usr/local/web_clients/web_client/dist"
const addressLogin = "/usr/local/web_clients/web_client_login/dist"

var addresses = AddressMap{
	webClient: address + webClient,
	webClientLogin: addressLogin + webClientLogin,
}
// Scripts -
func Scripts(w http.ResponseWriter, r *http.Request) {
	filePath, found := addresses[r.URL.Path]
	if !found {
		// 404
		return
	}
	http.ServeFile(w, r, filePath)
}
