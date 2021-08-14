package main

import (
	"io/ioutil"
	"time"

	"github.com/ravielze/oculi/generator/excel/excelize"
)

func main() {
	excel := excelize.New()
	style := excelize.NewStyle(
		excelize.FullBorder,
		excelize.Alignment{
			WrapText:   true,
			Horizontal: excelize.HorizontalAlignmentCenter,
			Vertical:   excelize.VerticalAlignmentCenter,
		})
	excel.NewSheet("Test")
	excel.DeleteSheet("Sheet1")
	excel.SetValue("Test", excelize.Cell{C: 0, R: 0}, "jason")
	excel.SetValue("Test", excelize.Cell{C: 0, R: 1}, 1)
	excel.SetValue("Test", excelize.Cell{C: 0, R: 2}, 5.02)
	excel.SetValue("Test", excelize.Cell{C: 0, R: 3}, map[string]interface{}{"A": 123, "B": "blabla"})
	excel.SetValues("Test", excelize.Cell{C: 5, R: 0}, []interface{}{
		"jason",
		"sayang",
		1,
		2,
		3,
		time.Now(),
	})
	excel.MergeCell("Test", excelize.Cell{C: 5, R: 3}.Axis(), excelize.Cell{C: 8, R: 3}.Axis())
	excel.CreateStyle("basic", style)
	excel.SetStyleRange("Test", "basic", excelize.Cell{C: 0, R: 0}, excelize.Cell{C: 1, R: 3})
	excel.SetStyle("Test", "basic", excelize.Cell{C: 5, R: 2})
	excel.SetStyleCol("Test", "basic", excelize.Cell{C: 10})
	excel.SetStyleColRange("Test", "basic", excelize.Cell{C: 6}, excelize.Cell{C: 8})
	buff, err := excel.Generate()
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("test.xlsx", buff.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
}
