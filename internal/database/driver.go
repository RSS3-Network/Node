package database

type Driver string

const (
	DriverPostgreSQL Driver = "postgres"
	DriverMySQL      Driver = "mysql"
)
