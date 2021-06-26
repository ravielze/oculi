package radix36

import (
	"math/big"

	"github.com/martinlindhe/base36"
)

func EncodeInt(value uint64) (string, error) {
	var i big.Int
	i.SetUint64(value)
	byteArr, err := i.GobEncode()
	if err != nil {
		return "", err
	}
	return base36.EncodeBytes(byteArr[1:]), nil
}

func DecodeBigInt(radix36 string) big.Int {
	var i big.Int
	byteArr := base36.DecodeToBytes(radix36)
	i.SetBytes(byteArr)
	return i
}
