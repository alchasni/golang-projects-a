package file

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jszwec/csvutil"
	"github.com/tealeg/xlsx/v3"
)

func BuildCSV(fPath string, data interface{}) (err error) {
	dataBytes, err := csvutil.Marshal(data)
	if err != nil {
		return
	}

	err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm)
	if err != nil {
		if err != os.ErrExist {
			return err
		}
		// TODO: log warning
	}

	err = ioutil.WriteFile(fPath, dataBytes, 0644)
	return
}

func BuildXLSX(fPath string, data interface{}) (err error) {
	dataBytes, err := csvutil.Marshal(data)
	if err != nil {
		return err
	}

	reader := csv.NewReader(bytes.NewReader(dataBytes))
	f := xlsx.NewFile()

	sheet, err := f.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	record, err := reader.Read()
	for err != nil {
		row := sheet.AddRow()
		row.SetHeight(14.0)

		for _, field := range record {
			cell := row.AddCell()
			cell.Value = field
		}

		record, err = reader.Read()
	}

	err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm)
	if err != nil {
		if err != os.ErrExist {
			return
		}
		// TODO: log warning
	}

	return f.Save(fPath)
}
