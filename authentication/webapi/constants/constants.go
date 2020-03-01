package constants

import "os"

// Stage -
const Stage = "STAGE"

// Production -
const Production = "PRODUCTION"

// Environment -
var Environment = os.Getenv(Stage)