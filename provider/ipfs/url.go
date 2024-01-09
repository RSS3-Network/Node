package ipfs

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ipfs/go-cid"
)

func ParseURL(contentURL string) (endpoint, path string, err error) {
	parsedURL, err := url.Parse(contentURL)
	if err != nil {
		return "", "", fmt.Errorf("parse content url: %w", err)
	}

	if parsedURL.Scheme == "ipfs" || parsedURL.Scheme == "ipns" {
		ipfsCid := parsedURL.Host
		urlPath := fmt.Sprintf("%s%s", parsedURL.Host, parsedURL.Path)

		if ipfsCid == "ipfs" {
			ipfsCid = strings.Split(parsedURL.Path, "/")[1]
			urlPath = parsedURL.Path[1:]
		}

		if _, err := cid.Parse(ipfsCid); err != nil {
			return "", "", fmt.Errorf("invalid content id %s: %w", ipfsCid, err)
		}

		return "", fmt.Sprintf("/%s/%s", parsedURL.Scheme, urlPath), nil
	}

	// Path gateway
	if strings.HasPrefix(parsedURL.Path, "/ipfs/") || strings.HasPrefix(parsedURL.Path, "/ipns/") {
		contentID := strings.TrimPrefix(strings.TrimPrefix(parsedURL.Path, "/ipfs/"), "/ipns/")

		if _, err := cid.Parse(contentID); err != nil {
			return "", "", fmt.Errorf("invalid content id %s: %w", contentID, err)
		}

		return parsedURL.Host, parsedURL.Path, nil
	}

	// Subdomain gateway
	if splits := strings.Split(parsedURL.Host, "."); len(splits) >= 4 && (splits[1] == "ipfs" || splits[1] == "ipns") {
		contentID := splits[0]

		if _, err := cid.Parse(contentID); err != nil {
			return "", "", fmt.Errorf("invalid content id %s: %w", contentID, err)
		}

		return strings.Join(splits[2:], "."), fmt.Sprintf("/%s/%s%s", splits[1], contentID, parsedURL.Path), nil
	}

	// DNSLink gateway is not supported

	return "", "", fmt.Errorf("invalid content url %s", contentURL)
}
