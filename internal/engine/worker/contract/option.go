package contract

type Option struct {
	ParseTokenMetadata bool `yaml:"parse_token_metadata"`
}

func NewOption(parseTokenMetadata bool) (*Option, error) {
	var instance Option

	instance.ParseTokenMetadata = parseTokenMetadata

	return &instance, nil
}
