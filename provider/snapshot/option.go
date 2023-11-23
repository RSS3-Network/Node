package snapshot

// Option is a function that can be used to configure a [client].
type Option func(client *client) error

// WithAPIKey is an [Option] that can be used to set the X-API-KEY header for
func WithAPIKey(apiKey string) Option {
	return func(client *client) error {
		client.apiKey = apiKey

		return nil
	}
}
