package radix36

import (
	"math/big"
	"strings"

	uuid "github.com/gofrs/uuid"

	"github.com/martinlindhe/base36"
)

func EncodeUUID4() (string, error) {
	var i big.Int
	uuid, erru := uuid.NewV4()
	if erru != nil {
		return "", erru
	}
	i.SetString(strings.Replace(uuid.String(), "-", "", 4), 16)
	byteArr, err := i.GobEncode()
	if err != nil {
		return "", err
	}
	return base36.EncodeBytes(byteArr[1:]), nil
}

func EncodeUUID(uuid string) (string, error) {
	var i big.Int
	i.SetString(strings.Replace(uuid, "-", "", 4), 16)
	byteArr, err := i.GobEncode()
	if err != nil {
		return "", err
	}
	return base36.EncodeBytes(byteArr[1:]), nil
}

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

func DecodeBytes(radix36 string) []byte {
	byteArr := base36.DecodeToBytes(radix36)
	return byteArr
}

func DecodeUUID(radix36 string) uuid.UUID {
	byteArr := base36.DecodeToBytes(radix36)
	return uuid.FromBytesOrNil(byteArr)
}
