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

	xlsx := excelize.NewFile()
	sheet := "Sheet1"
	index := xlsx.NewSheet(sheet)

	colNames := make([]string, len(header))
	for i := range header {
		colNames[i], _ = excelize.ColumnNumberToName(i + 1)
	}

	// Set first row as header
	for i, entry := range header {
		xlsx.SetCellStr(sheet, fmt.Sprintf("%s%d", colNames[i], 1), entry)
	}

rows:
	for row := 2; ; row++ {
		fields, err := r.Read()
		switch {
		case err == io.EOF:
			break rows
		case err != nil:
			log.Fatal(err)
		}

		var n int
		if len(fields) < len(header) {
			n = len(fields)
		} else {
			n = len(header)
		}

		var column int
		for ; column < n; column++ {
			xlsx.SetCellStr(
				sheet,
				fmt.Sprintf("%s%d", colNames[column], row),
				fields[column])
		}

		for ; column < len(header); column++ {
			xlsx.SetCellStr(
				sheet,
				fmt.Sprintf("%s%d", colNames[column], row),
				"")
		}
	}

	xlsx.SetActiveSheet(index)

	out := bufio.NewWriter(os.Stdout)
	_, err = xlsx.WriteTo(out)
	errCheck(err)
	errCheck(out.Flush())
}
