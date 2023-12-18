package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Direction --transform=snake --trimprefix=Direction --output direction_string.go --json
type Direction uint64

const (
	DirectionIn Direction = iota + 1
	DirectionOut
	DirectionSelf
)
