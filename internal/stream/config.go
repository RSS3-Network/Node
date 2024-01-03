package stream

type Config struct {
	Enable *bool  `mapstructure:"enable" validate:"required" default:"false"`
	Driver Driver `mapstructure:"driver" validate:"required" default:"kafka"`
	Topic  string `mapstructure:"topic" validate:"required" default:"rss3.node.feeds"`
	URI    string `mapstructure:"uri" validate:"required" default:"localhost:9092"`
}
