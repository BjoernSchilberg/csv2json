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

	"github.com/360EntSecGroup-Skylar/excelize"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createHeader(entry []map[string]interface{}) []string {
	names := make(map[string]string)
	for _, eintrag := range entry {
		for name := range eintrag {
			names[name] = name

		}
	}
	var t []string
	for k := range names {
		t = append(t, k)

	}
	sort.Strings(t)
	return t
}

func exportXLSX(records []map[string]interface{}, sheetName string) {

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(sheetName)

	if sheetName != "Sheet1" {
		xlsx.DeleteSheet("Sheet1")
	}

	var header []string
	header = createHeader(records)
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
