package arweave

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=Gateway --linecomment --output gateway_string.go
type Gateway uint

const (
	GatewayArweave    Gateway = iota // https://arweave.net/
	GatewayARIO                      // https://ar-io.dev/
	GatewayPermagate                 // https://permagate.io/
	GatewayLove4Src                  // https://love4src.com/
	GatewayARBR                      // https://arbr.pro/
	GatewayBobInstein                // https://bobinstein.com/
	GatewayAleko0o                   // https://aleko0o.store/
	GatewaySulapan                   // https://sulapan.com/
	GatewayFllstck                   // https://arweave.fllstck.dev/
	GatewayBicem                     // https://bicem.xyz
	GatewayDilsinay                  // https://dilsinay.online/
	GatewayLostGame                  // https://lostgame.online/
	GatewayKhalDrogo                 // https://khaldrogo.site/
	GatewayDasamuka                  // https://dasamuka.cloud/
	GatewayArendor                   // http://arendor.xyz/
	GatewayVelaryon                  // https://velaryon.xyz/
)

var DefaultGateways = []string{
	GatewayArweave.String(),
	GatewayARIO.String(),
	GatewayPermagate.String(),
	GatewayLove4Src.String(),
	GatewayARBR.String(),
}
