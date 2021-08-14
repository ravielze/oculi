package qrcode

type (
	QRCodeOption interface {
		Apply(qrcode *QRCode)
	}

	withSize               int
	withEncoding           struct{ Encoding }
	withErrCorrectionLevel struct{ ErrorCorrectionLevel }
)

func WithSize(size int) QRCodeOption {
	return withSize(size)
}

func WithEncoding(enc Encoding) QRCodeOption {
	return withEncoding{enc}
}

func WithErrCorrectionLevel(ecl ErrorCorrectionLevel) QRCodeOption {
	return withErrCorrectionLevel{ecl}
}

func (w withSize) Apply(qrcode *QRCode) {
	qrcode.Size = int(w)
}

func (w withEncoding) Apply(qrcode *QRCode) {
	qrcode.Encoding = w.Encoding
}

func (w withErrCorrectionLevel) Apply(qrcode *QRCode) {
	qrcode.ErrorCorrectionLevel = w.ErrorCorrectionLevel
}
