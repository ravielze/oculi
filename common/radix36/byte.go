package radix36

import "github.com/martinlindhe/base36"

func DecodeBytes(radix36 string) []byte {
	byteArr := base36.DecodeToBytes(radix36)
	return byteArr
}

func EncodeBytes(bytes []byte) string {
	return base36.EncodeBytes(bytes)
}
