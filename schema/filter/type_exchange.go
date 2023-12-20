package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=ExchangeType --transform=snake --trimprefix=TypeExchange --output type_exchange_string.go --json --sql
type ExchangeType uint64

//goland:noinspection GoMixedReceiverTypes
func (i ExchangeType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i ExchangeType) Tag() Tag {
	return TagExchange
}

const (
	TypeExchangeStaking ExchangeType = iota + 1
)
