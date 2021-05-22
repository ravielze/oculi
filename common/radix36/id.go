package radix36

import (
	"math/big"
	"strings"

	uuid "github.com/gofrs/uuid"
	"github.com/martinlindhe/base36"
)

var b36 = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z'}

func findBytes(f byte, l, r int) int {
	if r >= l {
		mid := l + (r-l)/2
		if b36[mid] == f {
			return mid
		} else if b36[mid] > f {
			return findBytes(f, l, mid-1)
		} else {
			return findBytes(f, mid+1, r)
		}
	}
	return -1
}

func ValidateRadix36(val string) bool {
	val = strings.ToUpper(val)
	for _, v := range []byte(val) {
		if idx := findBytes(v, 0, 35); idx == -1 {
			return false
		}
	}
	return true
}

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

func MustEncodeUUID(uuid string) string {
	var i big.Int
	i.SetString(strings.Replace(uuid, "-", "", 4), 16)
	byteArr, err := i.GobEncode()
	if err != nil {
		panic(err)
	}
	return base36.EncodeBytes(byteArr[1:])
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
