package generator_utils

import "strings"

type Placeholder struct {
	placeholder string
	value       string
}

func NewPlaceholder(placeholder, value string) Placeholder {
	return Placeholder{
		placeholder: placeholder,
		value:       value,
	}
}

func Replacer(raw string, placeholders []Placeholder) string {
	result := raw
	for _, p := range placeholders {
		result = strings.ReplaceAll(result, p.placeholder, p.value)
	}
	return result
}