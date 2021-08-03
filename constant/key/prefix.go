package consts

func ParameterPrefix(key string) string {
	return "parameter." + key
}

func QueryPrefix(key string) string {
	return "query." + key
}

func EchoPrefix(key string) string {
	return "echo." + key
}
