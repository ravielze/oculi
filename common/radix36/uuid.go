package radix36

import (
	"math/big"
	"strings"

	uuid "github.com/gofrs/uuid"
	"github.com/martinlindhe/base36"
)

func EncodeUUID4() (string, error) {
	uuid, erru := uuid.NewV4()
	if erru != nil {
		return "", erru
	}
	return EncodeUUID(uuid.String())
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

func DecodeUUID(radix36 string) uuid.UUID {
	byteArr := base36.DecodeToBytes(radix36)
	return uuid.FromBytesOrNil(byteArr)
}

func MustEncodeUUID(uuid string) string {
	var i big.Int
	i.SetString(strings.Replace(uuid, "-", "", 4), 16)
	byteArr, err := i.GobEncode()
	if err != nil {
		panic(err)
	}
	return base36.EncodeBytes(byteArr[1:])
}
