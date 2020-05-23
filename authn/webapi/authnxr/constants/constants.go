package constants

import (
	"os"
)

var (
	Environment = os.Getenv("STAGE")

	AuthnUrlAddress = os.Getenv("AUTHN_URL_ADDRESS")
)