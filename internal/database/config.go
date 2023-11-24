package database

type Config struct {
	Driver Driver `mapstructure:"driver"`
	Mode   Mode   `mapstructure:"mode"`
	URI    string `mapstructure:"uri"`
}
