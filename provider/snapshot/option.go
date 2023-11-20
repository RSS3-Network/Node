package snapshot

type Option func(client *client) error

func WithAPIKey(apiKey string) Option {
	return func(client *client) error {
		client.apiKey = apiKey

		return nil
	}
}
