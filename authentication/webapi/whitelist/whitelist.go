package whitelist

import (
	"fmt"

	"webapi/interfaces/whitelistx"
)

// HelloWorld - ping
func HelloWorld() {
	whitelistx.Get("YODAWG")
	fmt.Println("bridge the gap")
}


// verify token is in whitelist
// create token, add to whitelist, and return token
// remove token from whitelist

// store for jti, minimal security details
// jti: token_specific_secret_key

// Only required pulic methods
// Exists return Receipts
// - called every request

// Add - can only be called from our password / fingerprint
// authentication

// only two services allowed
// - nginx / initial loadbalancer
// - security logging (potentially)

// This is also where you'd log
// to a secondary service focused on
