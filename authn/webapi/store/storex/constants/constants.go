package constants

import "os"

const (
	pgHost         = "HOST_PGSQL"
	pgPort         = "PORT_PGSQL"
	pgUsername     = "POSTGRES_USER"
	pgPassword     = "POSTGRES_PASSWORD"
	pgDatabaseName = "POSTGRES_DB"
)

var (
	Host     = os.Getenv(pgHost)
	Port     = os.Getenv(pgPort)
	User     = os.Getenv(pgUsername)
	Password = os.Getenv(pgPassword)
	Database = os.Getenv(pgDatabaseName)
)
