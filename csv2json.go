// Copyright 2016 by Bjoern Schilberg
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)

	r := csv.NewReader(in)

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
		} else if err != nil {
			log.Fatal(err)
		}

		record := make(map[string]string)
		for i := range fields {
			if i == len(header) {
				break
			}

			record[header[i]] = fields[i]
		}

		records = append(records, record)
	}

	jsondata, err := json.Marshal(records) // convert to JSON

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsondata))
}
