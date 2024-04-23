package ethereum

import "time"

// DefaultCallOption is not enabled for retrying.
var DefaultCallOption = CallOption{
	RetryAttempts: 0,
}

type CallOption struct {
	RetryAttempts uint
	RetryDelay    time.Duration
}
