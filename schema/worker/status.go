package worker

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Status --linecomment --output status_string.go --json --yaml --sql
type Status uint64

const (
	Disabled  Status = iota // Disabled
	Indexing                // Indexing
	Ready                   // Ready
	Unhealthy               // Unhealthy
)
