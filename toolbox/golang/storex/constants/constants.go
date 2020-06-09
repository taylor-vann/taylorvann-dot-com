package constants

import "os"

const (
	envDatabaseName = "STORE_POSTGRES_DB"
	envUsername     = "STORE_POSTGRES_USER"
	envPassword     = "STORE_POSTGRES_PASSWORD"
	envHost         = "STORE_HOST_PGSQL"
	envPort         = "STORE_PORT_PGSQL"
)

var (
	Host     = os.Getenv(envHost)
	Port     = os.Getenv(envPort)
	User     = os.Getenv(envUsername)
	Password = os.Getenv(envPassword)
	Database = os.Getenv(envDatabaseName)
)
