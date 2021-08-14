package qrcode

import "github.com/boombuler/barcode/qr"

type (
	// The lower correction level, the less dense the qr code image is,
	// which improves minimum printing size.
	ErrorCorrectionLevel int
)

const (
	// Recovers 7% of data | Level L
	LowCorrection ErrorCorrectionLevel = iota
	// Recovers 15% of data | Level M
	MediumCorrection
	// Recovers 25% of data | Level Q
	MediumHighCorrection
	// Recovers 30% of data | Level H
	HighCorrection
)

var (
	correctionData = []qr.ErrorCorrectionLevel{
		qr.L, qr.M, qr.Q, qr.H,
	}
)

func (ecl ErrorCorrectionLevel) Convert() qr.ErrorCorrectionLevel {
	return correctionData[ecl]
}
