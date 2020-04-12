package constants

import "os"

const (
// Stage -
	Stage = "STAGE"

// Production -
	Production = "PRODUCTION"
	Development = "DEVELOPMENT"
	Local = "LOCAL"

	TaylorVannDotCom = "taylorvann.com"
)
// Environment -
var Environment = os.Getenv(Stage)