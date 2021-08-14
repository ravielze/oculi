package qrcode

import "github.com/boombuler/barcode/qr"

type (
	ErrorCorrectionLevel int
)

const (
	// Recovers 7% of data
	LowCorrection ErrorCorrectionLevel = iota
	// Recovers 15% of data
	MediumCorrection
	// Recovers 25% of data
	MediumHighCorrection
	// Recovers 30% of data
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
