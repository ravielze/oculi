package getOrDefault

import "time"

func Time(val *time.Time, def time.Time) time.Time {
	empty := time.Time{}
	if val == nil || (*val).IsZero() || *val == empty {
		return def
	}
	return *val
}
