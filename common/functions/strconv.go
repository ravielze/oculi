package functions

import "strconv"

// Ascii TO Integer
func Atoi(val string, defaultValue uint64) uint64 {
	v, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}
