package arweave

import (
	"bufio"
	"context"
	"net/http"

	"github.com/samber/lo"
)

var DefaultGateways = []string{
	EndpointArweave.String(),
	EndpointArweaveFllstck.String(),
	EndpointARIO.String(),
	EndpointPermagate.String(),
	EndpointLove4Src.String(),
}

const (
	DefaultGatewayList = "https://gist.githubusercontent.com/pseudoyu/39664a873726fa0ec03d9ad163ece124/raw/4f59d1899725606d18b94ab8053339974d1fc241/arweave_gateways.txt"
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
