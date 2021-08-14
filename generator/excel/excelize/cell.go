package excelize

import (
	"fmt"
)

type (
	Cell struct {
		// R for Row
		R uint
		// C for Column
		C uint
	}
)

func (c Cell) Col() string {
	col := ""

	i := c.C
	for {
		remainder := i % 26
		col = fmt.Sprintf("%c%s", int('A'+remainder), col)
		i = i / 26
		if i < 1 {
			break
		}
		i--
	}
	return col
}

func (c Cell) Axis() string {
	return fmt.Sprintf("%s%d", c.Col(), c.R+1)
}

func (from Cell) RangeCell(to Cell) string {
	if from.R > to.R {
		return to.RangeCell(from)
	}
	return fmt.Sprintf("%s:%s", from.Axis(), to.Axis())
}

func (from Cell) RangeColumn(to Cell) string {
	if from.C < to.C {
		return to.RangeColumn(from)
	}
	return fmt.Sprintf("%s:%s", from.Col(), to.Col())
}
