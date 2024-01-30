package constant

import "fmt"

const Name = "node"

// The Go compiler will set these variables via ldflags.
var (
	Version string
	Commit  string
)

func BuildVersion() string {
	if Version == "" {
		Version = "0.0.0"
	}

	if Commit == "" {
		Commit = "000000"
	}

	return fmt.Sprintf("%s (%s)", Version, Commit)
}
