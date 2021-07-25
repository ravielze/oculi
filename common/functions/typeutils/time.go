package typeutils

import "time"

func Time(val *time.Time, def time.Time) time.Time {
	empty := time.Time{}
	if val == nil || (*val).IsZero() || *val == empty {
		return def
	}
	return *val
}

func TimeOrNil(val *time.Time) *time.Time {
	empty := time.Time{}
	if val == nil || (*val).IsZero() || *val == empty {
		return nil
	}
	return val
}
