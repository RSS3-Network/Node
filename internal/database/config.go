package database

type Config struct {
	Driver    Driver `mapstructure:"driver" validate:"required" default:"cockroachdb"`
	Partition bool   `mapstructure:"partition" validate:"required" default:"true"`
	URI       string `mapstructure:"uri" validate:"required" default:"postgres://root@localhost:26257/defaultdb"`
}
