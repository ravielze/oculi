package excelize

import (
	"bytes"
	"io"
	"reflect"

	"github.com/gabriel-vasile/mimetype"
	consts "github.com/ravielze/oculi/constant/errors"
)

func (f *File) SetValue(sheetName string, loc Cell, value interface{}) error {
	if loc.R >= MaxRow || loc.C >= MaxCol {
		return consts.ErrCellOutOfRange
	}
	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Ptr {
		if !val.IsNil() {
			value = val.Elem()
		} else {
			value = ""
		}
	}
	return f.File.SetCellValue(sheetName, loc.Axis(), value)
}

func (f *File) SetValues(sheetName string, startCell Cell, value []interface{}) error {
	length := len(value)
	if length == 0 {
		return nil
	}
	if startCell.R >= MaxRow || startCell.C >= MaxCol {
		return consts.ErrCellOutOfRange
	}
	if startCell.C+uint(length) > MaxCol {
		return consts.ErrCellOutOfRange
	}

	for i := 0; i < length; i++ {
		if err := f.SetValue(sheetName, startCell, value[i]); err != nil {
			return err
		}
		startCell.C++
	}
	return nil
}

// Todo Merge Cell

func (f *File) SetImage(sheetName string, loc Cell, name string, data io.ReadSeeker) error {
	_, err := data.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	mime, err := mimetype.DetectReader(data)
	if err != nil {
		return err
	}

	_, err = data.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	_, err = io.Copy(&buff, data)
	if err != nil {
		return err
	}

	return f.File.AddPictureFromBytes(sheetName, loc.Axis(), "", name, mime.Extension(), buff.Bytes())
}
