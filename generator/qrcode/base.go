package qrcode

import (
	"bytes"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type (
	QRCode struct {
		Content              string
		ErrorCorrectionLevel ErrorCorrectionLevel
		Encoding             Encoding
		Size                 int
		data                 barcode.Barcode
	}
)

func New(content string, opts ...QRCodeOption) *QRCode {
	result := &QRCode{
		Content:              content,
		Size:                 200,
		Encoding:             Auto,
		ErrorCorrectionLevel: MediumHighCorrection,
		data:                 nil,
	}
	for _, opt := range opts {
		opt.Apply(result)
	}
	return result
}

func (q *QRCode) process() error {
	qrCode, err := qr.Encode(q.Content, q.ErrorCorrectionLevel.Convert(), q.Encoding.Convert())
	if err != nil {
		q.data = nil
		return err
	}

	resizedQRCode, err := barcode.Scale(qrCode, q.Size, q.Size)
	if err != nil {
		q.data = nil
		return err
	}

	q.data = resizedQRCode
	return nil
}

func (q *QRCode) FGenerate(fileName string) error {
	if q.data == nil {
		err := q.process()
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, q.data)
}

func (q *QRCode) Generate() (*bytes.Buffer, error) {
	if q.data == nil {
		err := q.process()
		if err != nil {
			return nil, err
		}
	}
	var result bytes.Buffer
	err := png.Encode(&result, q.data)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
