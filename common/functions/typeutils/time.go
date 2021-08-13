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

func TimePtr(val time.Time) *time.Time {
	i := val
	return &i
}

func Duration(val *time.Duration, def time.Duration) time.Duration {
	empty := time.Duration(0)
	if val == nil || *val == empty {
		return def
	}
	return *val
}

func DurationOrNil(val *time.Duration) *time.Duration {
	empty := time.Duration(0)
	if val == nil || *val == empty {
		return nil
	}
	return val
}

func DurationPtr(val time.Duration) *time.Duration {
	i := val
	return &i
}
