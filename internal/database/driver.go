package database

type Driver string

const (
	DriverPostgreSQL Driver = "postgresql"
	DriverMySQL      Driver = "mysql"
)
