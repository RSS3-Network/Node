package redis

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
)

// NewClient creates a new Redis client.
func NewClient(option config.Redis) (rueidis.Client, error) {
	clientOption := rueidis.ClientOption{
		InitAddress:  []string{option.Endpoint},
		Username:     option.Username,
		Password:     option.Password,
		DisableCache: option.DisableCache,
	}

	// Enable TLS if it is configured
	if option.TLS.Enabled {
		tlsConfig, err := buildTLSConfig(option.TLS)
		if err != nil {
			return nil, fmt.Errorf("failed to build TLS config: %w", err)
		}

		clientOption.TLSConfig = tlsConfig
	}

	return rueidis.NewClient(clientOption)
}

// buildTLSConfig builds a TLS configuration from the given TLS configuration.
func buildTLSConfig(tlsConfig config.RedisTLS) (*tls.Config, error) {
	tlsConf := &tls.Config{
		MinVersion: tls.VersionTLS12, // Set minimum TLS version to 1.2 to address G402 warning
	}

	if tlsConfig.CAFile != "" {
		caCert, err := os.ReadFile(tlsConfig.CAFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read CA certificate file: %w", err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConf.RootCAs = caCertPool
	}

	if tlsConfig.CertFile != "" && tlsConfig.KeyFile != "" {
		cert, err := tls.LoadX509KeyPair(tlsConfig.CertFile, tlsConfig.KeyFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load client certificate and key: %w", err)
		}

		tlsConf.Certificates = []tls.Certificate{cert}
	}

	tlsConf.InsecureSkipVerify = tlsConfig.InsecureSkipVerify

	return tlsConf, nil
}
