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
)

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

func exportCSV(lines []map[string]interface{}) {
	out := csv.NewWriter(bufio.NewWriter(os.Stdout))

	var header, record []string

	for line, entry := range lines {
		if header == nil {
			header = createHeader(entry)
			record = make([]string, len(header))
			if err := out.Write(header); err != nil {
				log.Fatalf("Error writing CSV: %v\n", err)
			}
		}
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

		exportCSV(target)

	} else {
		var lines []map[string]interface{}
		err := decoder.Decode(&lines)
		if err != nil {
			log.Fatalf("Cannot decode JSON file: %v.\n", err)
		}
		exportCSV(lines)
	}

}
