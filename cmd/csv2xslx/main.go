// Copyright 2019 by Bjoern Schilberg
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	r := csv.NewReader(bufio.NewReader(os.Stdin))
	r.FieldsPerRecord = -1

	header, err := r.Read()
	errCheck(err)

	var records []map[string]string

	for {
		fields, err := r.Read()
		if err == io.EOF {
			break
		}
		errCheck(err)

		record := map[string]string{}
		for i, name := range header {
			if i == len(fields) {
				break
			}
			record[name] = fields[i]
		}

		records = append(records, record)
	}

	xlsx := excelize.NewFile()
	sheet := "Sheet1"
	index := xlsx.NewSheet(sheet)

	// Set first row as header
	for cellIndex, entry := range header {
		col, _ := excelize.ColumnNumberToName(cellIndex + 1)
		xlsx.SetCellStr(sheet, fmt.Sprintf("%s%d", col, 1), entry)
	}

	for rowIndex, entry := range records {

		for cellIndex, name := range header {
			col, _ := excelize.ColumnNumberToName(cellIndex + 1)
			if _, found := entry[name]; found {
				xlsx.SetCellStr(sheet, fmt.Sprintf("%s%d", col, rowIndex+2), entry[name])
			} else {
				xlsx.SetCellStr(sheet, fmt.Sprintf("%s%d", col, rowIndex+2), "")
			}
			cellIndex++
		}
	}
	xlsx.SetActiveSheet(index)

	buf, _ := xlsx.WriteToBuffer()
	_, errWrite := buf.WriteTo(os.Stdout)
	errCheck(errWrite)

	//errCheck(xlsx.SaveAs("out.xlsx"))

}
