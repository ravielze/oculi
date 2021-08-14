package excelize

import (
	"github.com/danhper/structomap"
)

type (
	UnderlineType int

	Font struct {
		Bold      bool
		Italic    bool
		Underline UnderlineType
		Family    string
		Size      float64
		Strike    bool
		Color     string
	}
)

const (
	UnderlineNone UnderlineType = iota
	UnderlineSingle
	UnderlineDouble
)

var underlineTypeData = []string{
	"",
	"single",
	"double",
}

func (u UnderlineType) String() string {
	return underlineTypeData[u]
}

func (f Font) Style() map[string]interface{} {
	result := structomap.New().UseSnakeCase().
		PickAll().
		PickFunc(func(i interface{}) interface{} {
			return i.(UnderlineType).String()
		}, "Underline").
		PickFunc(func(i interface{}) interface{} {
			data := i.(string)
			if data == "" {
				return Black
			}
			return data
		}, "Color").
		OmitIf(func(i interface{}) bool {
			data := i.(Font).Family
			return data == ""
		}, "Family").
		OmitIf(func(i interface{}) bool {
			data := i.(Font).Underline
			return data == UnderlineNone
		}, "Underline").
		Transform(f)
	return result
}

func (w Font) Apply(s Style) {
	s.data["font"] = w.Style()
}
