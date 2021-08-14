package qrcode

import "github.com/boombuler/barcode/qr"

type (
	Encoding int
)

const (
	// Auto will choose ths best matching encoding
	Auto Encoding = iota
	// Numeric encoding only encodes numbers [0-9]
	Numeric
	// AlphaNumeric encoding only encodes uppercase letters, numbers and  [Space], $, %, *, +, -, ., /, :
	AlphaNumeric
	// Unicode encoding encodes the string as utf-8
	Unicode
)

var (
	encodingData = []qr.Encoding{
		qr.Auto, qr.Numeric, qr.AlphaNumeric, qr.Unicode,
	}
)

func (enc Encoding) Convert() qr.Encoding {
	return encodingData[enc]
}
