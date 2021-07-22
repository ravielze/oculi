package getOrDefault

func String(val *string, def string) string {
	if val == nil || *val == "" {
		return def
	}
	return *val
}
