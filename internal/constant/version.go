package constant

import "fmt"

var (
	ServiceVersion string
	ServiceCommit  string
)

func BuildServiceVersion() string {
	if ServiceVersion == "" {
		ServiceVersion = "0.0.0"
	}

	if ServiceCommit == "" {
		ServiceCommit = "000000"
	}

	return fmt.Sprintf("%s (%s)", ServiceVersion, ServiceCommit)
}
