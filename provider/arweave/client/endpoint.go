package client

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=Endpoint --linecomment --output endpoint_string.go
type Endpoint int

const (
	EndpointArweave        Endpoint = iota + 1 // https://arweave.net/
	EndpointArweaveFllstck                     // https://arweave.fllstck.dev/
	EndpointARIO                               // https://ar-io.dev/
	EndpointPermagate                          // https://permagate.io/
	EndpointLove4Src                           // https://love4src.com/
)
