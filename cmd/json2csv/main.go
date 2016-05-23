// Copyright 2016 by Bjoern Schilberg
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
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

func main() {

	decoder := json.NewDecoder(bufio.NewReader(os.Stdin))

	var lines []map[string]interface{}
	err := decoder.Decode(&lines)
	if err != nil {
		log.Fatalf("Cannot decode JSON file: %v.\n", err)
	}

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
				value = fmt.Sprintf("%s", v)
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
