package constants

import "os"

const (
	Stage = "STAGE"

	Production = "PRODUCTION"
	Development = "DEVELOPMENT"
	Local = "LOCAL"

	TaylorVannDotCom = "taylorvann.com"
)

var Environment = os.Getenv(Stage)