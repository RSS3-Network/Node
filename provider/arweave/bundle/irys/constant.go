package irys

//go:generate go run --mod=mod github.com/Khan/genqlient@v0.7.0

const (
	GatewayMainnet = "https://gateway.irys.xyz"

	EndpointMainnet = "https://arweave.mainnet.irys.xyz/graphql"
	EndpointDevnet  = "https://arweave.devnet.irys.xyz/graphql"

	// DefaultLimit is the limit on the number of items when obtaining transactions, with a maximum value of 1000.
	DefaultLimit = 1000
)

// DefaultGateways is the list of default gateways.
var DefaultGateways = []string{
	GatewayMainnet,
}
