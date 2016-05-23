// Copyright 2016 by Bjoern Schilberg
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {

	r := csv.NewReader(bufio.NewReader(os.Stdin))
	r.FieldsPerRecord = -1

	header, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	var records []map[string]string

	for {
		fields, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		record := map[string]string{}
		for i, name := range header {
			if i == len(fields) {
				break
			}
			record[name] = fields[i]
		}

		records = append(records, record)
	}

	encoder := json.NewEncoder(bufio.NewWriter(os.Stdout))
	if err := encoder.Encode(records); err != nil {
		log.Fatal(err)
	}
}
