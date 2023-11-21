package database

type Driver string

const (
	DriverCockroach Driver = "cockroach"
	DriverPostgres  Driver = "postgres"
	DriverMySQL     Driver = "mysql"
)
