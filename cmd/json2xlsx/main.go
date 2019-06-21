// Copyright 2016 by Bjoern Schilberg
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createHeader(entry map[string]interface{}) []string {
	names := make([]string, len(entry))
	var i int
	for name := range entry {
		names[i] = name
		i++
	}
	sort.Strings(names)
	return names
}

func exportXLSX(records []map[string]interface{}, sheetName string) {

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(sheetName)

	if sheetName != "Sheet1" {
		xlsx.DeleteSheet("Sheet1")
	}

	var header []string

	for rowIndex, entry := range records {
		header = createHeader(entry)
		colNames := make([]string, len(header))
		for i := range header {
			colNames[i], _ = excelize.ColumnNumberToName(i + 1)
		}

		// Set first row as header
		for i, entry := range header {
			xlsx.SetCellStr(sheetName, colNames[i]+"1", entry)
		}

		for cellIndex, name := range header {
			col, _ := excelize.ColumnNumberToName(cellIndex + 1)
			if _, found := entry[name]; found {
				switch t := entry[name].(type) {
				case string:
					xlsx.SetCellStr(sheetName, fmt.Sprintf("%s%d", col, rowIndex+2), t)
				case float64:
					xlsx.SetCellStr(sheetName, fmt.Sprintf("%s%d", col, rowIndex+2), fmt.Sprint(t))
				}
			} else {
				xlsx.SetCellStr(sheetName, fmt.Sprintf("%s%d", col, rowIndex+2), "")
			}
			cellIndex++
		}
	}
	xlsx.SetActiveSheet(index)

	buf, _ := xlsx.WriteToBuffer()
	_, errWrite := buf.WriteTo(os.Stdout)
	errCheck(errWrite)

}

func main() {

	var sheetName string
	var arrayStructure bool
	var objectName string

	flag.StringVar(&sheetName, "sheet", "Sheet1", "Set the worksheet name")
	flag.StringVar(&sheetName, "s", "Sheet1", "Set the worksheet name (shorthand).")
	flag.BoolVar(&arrayStructure, "array", true, "Is pure JSON array structure")
	flag.BoolVar(&arrayStructure, "a", true, "Is pure JSON array structure (shorthand).")
	flag.StringVar(&objectName, "object", "", "The name of the JSON object that holds the JSON array structure")
	flag.StringVar(&objectName, "o", "", "The name of the JSON object that holds the JSON array structure (shorthand).")

	flag.Parse()

	decoder := json.NewDecoder(bufio.NewReader(os.Stdin))

	if !arrayStructure {
		var obj map[string]interface{}
		err := decoder.Decode(&obj)
		if err != nil {
			log.Fatalf("Cannot decode JSON file: %v.\n", err)
		}

		var target []map[string]interface{}
		for _, v := range obj[objectName].([]interface{}) {
			target = append(target, v.(map[string]interface{}))

		}

		exportXLSX(target, sheetName)

	} else {
		var records []map[string]interface{}
		err := decoder.Decode(&records)
		if err != nil {
			log.Fatalf("Cannot decode JSON file: %v.\n", err)
		}
		exportXLSX(records, sheetName)
	}

}
