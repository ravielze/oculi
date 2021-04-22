package random

import (
	"math/big"
	"strings"

	"github.com/martinlindhe/base36"
	uuid "github.com/satori/go.uuid"
)

func NewUUID4ToRadix36() (string, error) {
	var i big.Int
	uuid := uuid.NewV4().String()
	i.SetString(strings.Replace(uuid, "-", "", 4), 16)
	byteArr, err := i.GobEncode()
	if err != nil {
		return "", err
	}
	return base36.EncodeBytes(byteArr[1:]), nil
}
