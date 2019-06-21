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
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createHeader(entries []map[string]interface{}) []string {
	uniqueMap := make(map[string]string)
	for _, entry := range entries {
		for name := range entry {
			uniqueMap[name] = ""

		}
	}
	var names []string
	for k := range uniqueMap {
		names = append(names, k)

	}
	sort.Strings(names)
	return names
}

func exportXLSX(records []map[string]interface{}, sheetName string, columns []string) {

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(sheetName)

	if sheetName != "Sheet1" {
		xlsx.DeleteSheet("Sheet1")
	}

	var header []string

	if header == nil && len(columns) <= 0 {
		header = createHeader(records)
	} else {
		header = columns
	}
	colNames := make([]string, len(header))
	for i := range header {
		colNames[i], _ = excelize.ColumnNumberToName(i + 1)
	}
	// Set first row as header
	for i, entry := range header {
		xlsx.SetCellStr(sheetName, colNames[i]+"1", entry)
	}

	for rowIndex, entry := range records {
		rowS := strconv.Itoa(rowIndex + 2)
		for cellIndex, name := range header {
			if _, found := entry[name]; found {
				switch t := entry[name].(type) {
				case string:
					xlsx.SetCellStr(sheetName, colNames[cellIndex]+rowS, t)
				case float64:
					xlsx.SetCellStr(sheetName, colNames[cellIndex]+rowS, fmt.Sprint(t))
				}
			} else {
				xlsx.SetCellStr(sheetName, colNames[cellIndex]+rowS, "")
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
	var listOfColumns string
	var columns []string

	flag.StringVar(&sheetName, "sheet", "Sheet1", "Set the worksheet name")
	flag.StringVar(&sheetName, "s", "Sheet1", "Set the worksheet name (shorthand).")
	flag.BoolVar(&arrayStructure, "array", true, "Is pure JSON array structure")
	flag.BoolVar(&arrayStructure, "a", true, "Is pure JSON array structure (shorthand).")
	flag.StringVar(&objectName, "object", "", "The name of the JSON object that holds the JSON array structure")
	flag.StringVar(&objectName, "o", "", "The name of the JSON object that holds the JSON array structure (shorthand).")
	flag.StringVar(&listOfColumns, "c", "", "List of columns which should be exported (shorthand).")
	flag.StringVar(&listOfColumns, "columns", "", "List of columns which should be exported (shorthand).")

	flag.Parse()

	if len(listOfColumns) > 0 {
		columns = strings.Split(listOfColumns, ",")
	}

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

		exportXLSX(target, sheetName, columns)

	} else {
		var records []map[string]interface{}
		err := decoder.Decode(&records)
		if err != nil {
			log.Fatalf("Cannot decode JSON file: %v.\n", err)
		}
		exportXLSX(records, sheetName, columns)
	}

}
