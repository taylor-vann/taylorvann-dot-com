// brian taylor vann
// briantaylorvann dot com

package constants

import "os"

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
