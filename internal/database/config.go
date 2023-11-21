package database

type Config struct {
	Driver Driver `mapstructure:"driver"`
	URI    string `mapstructure:"uri"`
}
