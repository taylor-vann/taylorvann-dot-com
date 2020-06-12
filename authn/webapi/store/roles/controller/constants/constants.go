// brian taylor vann
// briantaylorvann dot com

package constants

import "os"

const (
	Production 			= "PRODUCTION"
	Development			= "DEVELOPMENT"
	Local						= "LOCAL"

	Roles     			= "roles"
	RolesTest 			= "roles_test"
	RolesUnitTests  = "roles_unit_tests"
)

var (
	Environment = os.Getenv("STAGE")
)

