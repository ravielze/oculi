package excelize

import (
	"github.com/danhper/structomap"
)

type (
	BorderStyle int
	BorderType  int

	Border struct {
		Stylee BorderStyle
		Type   BorderType
		Color  string
	}

	BorderGroup struct {
		Stylee BorderStyle
		Color  string
	}
)

const (
	BorderNone BorderStyle = iota
	BorderLine
	BorderLineSemiBold
	BorderDash
	BorderDot
	BorderLineBold
	BorderDoubleLine
	BorderDotDense
	BorderDashBold
	BorderDashDot
	BorderDashDotBold
	BorderDashDotDot
	BorderDashDotDotBold
	BorderSlantDashDot
)

const (
	BorderTypeLeft BorderType = iota
	BorderTypeRight
	BorderTypeTop
	BorderTypeBottom
	BorderTypeDiagonalUp
	BorderTypeDiagonalDown
)

var (
	FullBorder StyleableOption = BorderGroup{Stylee: BorderLine}
)

var borderTypeData = []string{
	"left",
	"right",
	"top",
	"bottom",
	"diagonalUp",
	"diagonalDown",
}

func (b BorderType) String() string {
	return borderTypeData[b]
}

func (b Border) Style() map[string]interface{} {
	result := structomap.New().UseSnakeCase().
		PickFunc(func(i interface{}) interface{} {
			data := i.(string)
			if data == "" {
				return Black
			}
			return data
		}, "Color").
		AddFunc("Style", func(i interface{}) interface{} {
			return i.(Border).Stylee
		}).
		PickFunc(func(i interface{}) interface{} {
			return i.(BorderType).String()
		}, "Type").
		Transform(b)
	return result
}

func (w BorderGroup) Apply(s Style) {
	for i := 0; i < 4; i++ {
		Border{
			Stylee: w.Stylee,
			Color:  w.Color,
			Type:   BorderType(i),
		}.Apply(s)
	}
}

func (w Border) Apply(s Style) {
	if w.Stylee == 0 {
		return
	}
	if data, ok := s.data["border"]; ok {
		borders := data.([]map[string]interface{})
		borders = append(borders, w.Style())
		s.data["border"] = borders
		return
	}
	s.data["border"] = []map[string]interface{}{w.Style()}
}
