package constants

import "os"

type TableNames struct {
	Users string
}

const (
	Production 			= "PRODUCTION"
	Development			= "DEVELOPMENT"
	Local						= "LOCAL"

	Users     			= "users"
	UsersTest 			= "users_test"
	UsersUnitTests  = "users_unit_tests"
)

var (
	Environment = os.Getenv("STAGE")
)
