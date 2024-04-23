package ethereum

import "time"

var DefaultCallOption = CallOption{
	RetryAttempts: 0,
}

type CallOption struct {
	RetryAttempts uint
	RetryDelay    time.Duration
}
