package core

import "time"

type Metrics struct {
	Comparisons int
	Swaps       int
	Time        time.Duration
	Memory      int64
}
