package typeutils

func Int(val *int, def int) int {
	if val == nil || *val == 0 {
		return def
	}
	return *val
}

func IntOrNil(val *int) *int {
	if val == nil || *val == 0 {
		return nil
	}
	return val
}

func Uint(val *uint, def uint) uint {
	if val == nil || *val == uint(0) {
		return def
	}
	return *val
}

func UintOrNil(val *uint) *uint {
	if val == nil || *val == uint(0) {
		return nil
	}
	return val
}

func Int64(val *int64, def int64) int64 {
	if val == nil || *val == int64(0) {
		return def
	}
	return *val
}

func Int64OrNil(val *int64) *int64 {
	if val == nil || *val == int64(0) {
		return nil
	}
	return val
}

func Uint64(val *uint64, def uint64) uint64 {
	if val == nil || *val == uint64(0) {
		return def
	}
	return *val
}

func Uint64OrNil(val *uint64) *uint64 {
	if val == nil || *val == uint64(0) {
		return nil
	}
	return val
}
