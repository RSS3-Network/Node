package ai

import (
	"github.com/rss3-network/node/v2/config"
)

type Option struct {
	AgentdataDBURL string        `json:"agentdata_db_url" mapstructure:"agentdata_db_url"`
	OpenAIAPIKey   string        `json:"openai_api_key" mapstructure:"openai_api_key"`
	OllamaHost     string        `json:"ollama_host" mapstructure:"ollama_host"`
	Twitter        OptionTwitter `json:"twitter" mapstructure:"twitter"`
	KaitoAPIToken  string        `json:"kaito_api_token" mapstructure:"kaito_api_token"`
}

type OptionTwitter struct {
	BearerToken       string `json:"bearer_token" mapstructure:"bearer_token"`
	APIKey            string `json:"api_key" mapstructure:"api_key"`
	APISecret         string `json:"api_secret" mapstructure:"api_secret"`
	AccessToken       string `json:"access_token" mapstructure:"access_token"`
	AccessTokenSecret string `json:"access_token_secret" mapstructure:"access_token_secret"`
}

func NewOption(options *config.Parameters) (*Option, error) {
	var instance Option

	if options == nil {
		return &instance, nil
	}

	if err := options.Decode(&instance); err != nil {
		return nil, err
	}

	return &instance, nil
}
