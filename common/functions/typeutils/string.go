package typeutils

func String(val *string, def string) string {
	if val == nil || *val == "" {
		return def
	}
	return *val
}

func StringOrNil(val *string) *string {
	if val == nil || *val == "" {
		return nil
	}
	return val
}
