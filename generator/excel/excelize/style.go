package excelize

import (
	consts "github.com/ravielze/oculi/constant/errors"
)

func (f *File) CreateStyle(styleName string, style Style) error {
	data, err := f.enc.Marshal(style.Style())
	if err != nil {
		return err
	}
	id, err := f.File.NewStyle(string(data))
	if err != nil {
		return err
	}
	f.styles[styleName] = id
	return nil
}

func (f *File) Style(styleName string) (int, error) {
	if data, ok := f.styles[styleName]; ok {
		return data, nil
	}
	return 0, consts.ErrStyleNotFound
}

func (f *File) SetStyle(sheetName, styleName string, loc Cell) error {
	styleId, err := f.Style(styleName)
	if err != nil {
		return err
	}
	err = f.File.SetCellStyle(sheetName, loc.Axis(), loc.Axis(), styleId)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) SetStyleRange(sheetName, styleName string, from, to Cell) error {
	styleId, err := f.Style(styleName)
	if err != nil {
		return err
	}
	err = f.File.SetCellStyle(sheetName, from.Axis(), to.Axis(), styleId)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) SetStyleCol(sheetName, styleName string, col Cell) error {
	styleId, err := f.Style(styleName)
	if err != nil {
		return err
	}
	err = f.File.SetColStyle(sheetName, col.Col(), styleId)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) SetStyleColRange(sheetName, styleName string, from, to Cell) error {
	styleId, err := f.Style(styleName)
	if err != nil {
		return err
	}
	err = f.File.SetColStyle(sheetName, from.RangeColumn(to), styleId)
	if err != nil {
		return err
	}
	return nil
}
