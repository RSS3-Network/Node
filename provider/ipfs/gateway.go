package ipfs

import (
	"bufio"
	"context"
	"net/http"

	"github.com/samber/lo"
)

const (
	DrfaultGatewayRSS3       = "https://ipfs.rss3.page/"
	DefaultGatewayIPFS       = "https://ipfs.io/"
	DefaultGatewayCloudflare = "https://cloudflare-ipfs.com/"
	DefaultGateway4EVERLAND  = "https://4everland.io/"
)

var DefaultGateways = []string{
	DrfaultGatewayRSS3,
	DefaultGatewayIPFS,
	DefaultGatewayCloudflare,
	DefaultGateway4EVERLAND,
}

const (
	DefaultGatewayList = "https://raw.githubusercontent.com/ipfs/public-gateway-checker/master/gateways.txt"
)

func FetchGateways(ctx context.Context, gatewayList string) ([]string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, gatewayList, nil)
	if err != nil {
		return nil, err
	}

	// nolint:bodyclose // False positive.
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer lo.Try(response.Body.Close)

	var (
		scanner     = bufio.NewScanner(response.Body)
		gatewayURLs = make([]string, 0)
	)

	for scanner.Scan() {
		gatewayURLs = append(gatewayURLs, scanner.Text())
	}

	return gatewayURLs, nil
}
