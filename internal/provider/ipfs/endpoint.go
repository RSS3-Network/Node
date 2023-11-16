package ipfs

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/samber/lo"
)

const (
	publicGatewaysURL = "https://raw.githubusercontent.com/ipfs/public-gateway-checker/master/gateways.txt"
	checkDefaultPath  = "/ipfs/bafybeibwzifw52ttrkqlikfzext5akxu7lz4xiwjgwzmqcpdzmp3n5vnbe"

	DefaultIPFSGateway = "https://gateway.ipfs.io"
	DefaultIPFSIO      = "https://ipfs.io"
	DefaultCloudflare  = "https://cloudflare-ipfs.com"
	Default4EVERLAND   = "https://4everland.io"
)

// GetIPFSPublicGateways gets a list of public IPFS gateways.
func GetIPFSPublicGateways(_ context.Context) ([]string, error) {
	DefaultGateways := []string{DefaultIPFSGateway, DefaultIPFSIO, DefaultCloudflare, Default4EVERLAND}

	// get public gateways from GitHub
	response, err := http.Get(publicGatewaysURL)
	if err != nil {
		return nil, fmt.Errorf("get public gateways error: %w", err)
	}

	defer lo.Try(response.Body.Close)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return DefaultGateways, nil
	}

	return lo.WithoutEmpty(strings.Split(string(data), "\n")), nil
}

// CheckIPFSEndpoints checks IPFS endpoints and sorts them by response time.
func CheckIPFSEndpoints(ctx context.Context, client http.Client, endpoints []string) (checkedEndpoints []string) {
	type record struct {
		endpoint string
		elapsed  int64
	}

	recordChan := make(chan record, len(endpoints))

	// invalidEndpoint handle invalid endpoint
	invalidEndpoint := func(endpoint string) {
		recordChan <- record{endpoint: endpoint, elapsed: math.MaxInt64}
	}

	var wg sync.WaitGroup

	for _, endpoint := range endpoints {
		endpoint := endpoint

		wg.Add(1)
		go func() {
			defer wg.Done()

			endpointURL, err := url.Parse(endpoint)
			if err != nil {
				invalidEndpoint(endpoint)

				return
			}

			endpointURL.Path = checkDefaultPath

			startTime := time.Now()

			request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL.String(), nil)
			if err != nil {
				invalidEndpoint(endpoint)

				return
			}

			response, err := client.Do(request)
			if err != nil {
				invalidEndpoint(endpoint)

				return
			}

			defer lo.Try(response.Body.Close)

			if response.StatusCode != http.StatusOK {
				invalidEndpoint(endpoint)

				return
			}

			recordChan <- record{endpoint: endpoint, elapsed: time.Since(startTime).Milliseconds()}
		}()
	}

	go func() {
		wg.Wait()
		close(recordChan)
	}()

	var results []record

	for record := range recordChan {
		results = append(results, record)
	}

	// sort by response time
	sort.Slice(results, func(i, j int) bool {
		return results[i].elapsed < results[j].elapsed
	})

	for _, result := range results {
		checkedEndpoints = append(checkedEndpoints, result.endpoint)
	}

	return checkedEndpoints
}
