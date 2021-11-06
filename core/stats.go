package core

import (
	"time"
)

type Stats struct {
	StartedAt  time.Time
	FinishedAt time.Time
}

type PingStats struct {
	ok               bool
	pingStartAt      time.Time
	pingFinishedAt   time.Time
	serverReceivedAt string
}
