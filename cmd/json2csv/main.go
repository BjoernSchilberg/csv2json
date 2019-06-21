// Copyright 2016 by Bjoern Schilberg
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

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

func exportCSV(lines []map[string]interface{}, columns []string) {
	out := csv.NewWriter(bufio.NewWriter(os.Stdout))

	var header, record []string

	if header == nil && len(columns) <= 0 {
		header = createHeader(lines)
	} else {
		header = columns
	}

	record = make([]string, len(header))
	if err := out.Write(header); err != nil {
		log.Fatalf("Error writing CSV: %v\n", err)
	}

	for line, entry := range lines {
		for i, name := range header {
			var value string
			if v, found := entry[name]; found {
				switch t := v.(type) {
				case string:
					value = fmt.Sprintf("%s", t)
				case float64:
					value = fmt.Sprint(t)
				}
			} else {
				log.Printf("key %v not found in line %d.\n", name, line+1)
			}
			record[i] = value
		}
		if err := out.Write(record); err != nil {
			log.Fatalf("Error writing CSV: %v\n", err)
		}
	}
	out.Flush()
	if err := out.Error(); err != nil {
		log.Fatalf("Error flushing CSV: %v\n", err)
	}

}

func main() {

	var arrayStructure bool
	var objectName string
	var listOfColumns string
	var columns []string

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

		exportCSV(target, columns)

	} else {
		var lines []map[string]interface{}
		err := decoder.Decode(&lines)
		if err != nil {
			log.Fatalf("Cannot decode JSON file: %v.\n", err)
		}
		exportCSV(lines, columns)
	}

}
