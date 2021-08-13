package typeutils

func BoolPtr(val bool) *bool {
	i := val
	return &i
}
