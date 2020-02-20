package constants

import "os"

// Stage -
const Stage = "STAGE"

// Production -
const Production = "PRODUCTION"

// TaylorVannDotCom -
const TaylorVannDotCom = "taylorvann.com"

// Environment -
var Environment = os.Getenv(Stage)