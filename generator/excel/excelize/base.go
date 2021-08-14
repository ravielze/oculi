package excelize

import (
	"bytes"

	"github.com/ravielze/oculi/encoding"
	"github.com/ravielze/oculi/encoding/jsoniter"
	"github.com/xuri/excelize/v2"
)

type (
	File struct {
		*excelize.File
		enc    encoding.Encoding
		styles map[string]int
	}

	Styleable interface {
		Style() map[string]interface{}
	}

	StyleableOption interface {
		Apply(style Style)
	}

	Style struct {
		data map[string]interface{}
	}
)

const (
	Black  = "#000000"
	White  = "#FFFFFF"
	MaxRow = excelize.TotalRows
	MaxCol = excelize.TotalColumns
)

// Implement Table
var (
	// To make sure these struct implement Interface
	// Interface      = Struct

	_ Styleable       = Style{}
	_ Styleable       = Alignment{}
	_ Styleable       = Border{}
	_ StyleableOption = Border{}
	_ StyleableOption = BorderGroup{}
	_ StyleableOption = Alignment{}
	_ StyleableOption = Font{}
)

func (s Style) Style() map[string]interface{} {
	return s.data
}

func NewStyle(opts ...StyleableOption) Style {
	style := Style{data: make(map[string]interface{})}
	for _, opt := range opts {
		opt.Apply(style)
	}
	return style
}

func New() *File {
	file := &File{
		File:   excelize.NewFile(),
		enc:    jsoniter.New(),
		styles: make(map[string]int),
	}
	return file
}

func (f *File) Generate() (*bytes.Buffer, error) {
	return f.File.WriteToBuffer()
}

func (f *File) FGenerate(fileName string) error {
	return f.File.SaveAs(fileName)
}
