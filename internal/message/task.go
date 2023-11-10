package message

import "time"

type Task interface {
	ID() string
	Timestamp() time.Time
}
