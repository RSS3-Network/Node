package engine

type Config struct {
	Network  string `mapstructure:"network"`
	Chain    string `mapstructure:"chain"`
	Endpoint string `mapstructure:"endpoint"`
}
