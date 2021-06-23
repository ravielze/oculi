package oculi_time

import "time"

var (
	isFreeze  = false
	savedTime *time.Time
)

func Now() time.Time {
	if isFreeze && savedTime != nil {
		return *savedTime
	}
	return time.Now()
}

func Mock(t time.Time) {
	isFreeze = true
	savedTime = &t
}

func Reset() {
	isFreeze = false
	savedTime = nil
}
