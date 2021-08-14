package excelize

import "github.com/danhper/structomap"

type (
	HorizontalAlignmentStyle int
	VerticalAlignmentStyle   int

	Alignment struct {
		Indent          int
		RelativeIndent  int
		ReadingOrder    int
		TextRotation    int
		ShrinkToFit     bool
		Horizontal      HorizontalAlignmentStyle
		Vertical        VerticalAlignmentStyle
		JustifyLastLine bool
		WrapText        bool
	}
)

const (
	HorizontalAlignmentDefault HorizontalAlignmentStyle = iota
	HorizontalAlignmentLeft
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
	HorizontalAlignmentFill
	HorizontalAlignmentJustify
	HorizontalAlignmentCenterContinues
	HorizontalAlignmentDistributed
)

const (
	VerticalAlignmentDefault VerticalAlignmentStyle = iota
	VerticalAlignmentTop
	VerticalAlignmentCenter
	VerticalAlignmentJustify
	VerticalAlignmentDistributed
)

var (
	horizontalAlignmentStyleData = []string{
		"left",
		"left",
		"center",
		"right",
		"fill",
		"justify",
		"centerContinuous",
		"distributed",
	}
	verticalAlignmentStyleData = []string{
		"top",
		"top",
		"center",
		"justify",
		"distributed",
	}
)

func (h HorizontalAlignmentStyle) String() string {
	return horizontalAlignmentStyleData[h]
}

func (v VerticalAlignmentStyle) String() string {
	return verticalAlignmentStyleData[v]
}

func (a Alignment) Style() map[string]interface{} {
	result := structomap.New().UseSnakeCase().
		PickAll().
		Omit("Indent").
		AddFunc("Ident", func(i interface{}) interface{} {
			return i.(Alignment).Indent
		}).
		PickFunc(func(i interface{}) interface{} {
			return i.(HorizontalAlignmentStyle).String()
		}, "Horizontal").
		PickFunc(func(i interface{}) interface{} {
			return i.(VerticalAlignmentStyle).String()
		}, "Vertical").
		Transform(a)
	return result
}

var emptyAlignment = Alignment{}

func (w Alignment) Apply(s Style) {
	if w == emptyAlignment {
		return
	}
	s.data["alignment"] = w.Style()
}
