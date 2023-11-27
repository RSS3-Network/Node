package database

type Config struct {
	Driver    Driver `mapstructure:"driver"`
	Partition bool   `mapstructure:"partition"`
	URI       string `mapstructure:"uri"`
}
