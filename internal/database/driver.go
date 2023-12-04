package database

type Driver string

const (
	DriverCockroachDB Driver = "cockroachdb"
	DriverPostgreSQL  Driver = "postgresql"
	DriverMySQL       Driver = "mysql"
)
