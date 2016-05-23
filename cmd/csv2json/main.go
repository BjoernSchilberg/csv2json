// Copyright 2016 by Bjoern Schilberg
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var pretty bool

	flag.BoolVar(&pretty, "pretty", false, "pretty print the JSON output.")
	flag.BoolVar(&pretty, "p", false, "pretty print the JSON output (shorthand).")

	flag.Parse()

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

	if pretty {
		b, err := json.Marshal(records)
		errCheck(err)
		var out bytes.Buffer
		json.Indent(&out, b, "", "\t")
		_, err = out.WriteTo(os.Stdout)
		errCheck(err)
	} else {
		encoder := json.NewEncoder(bufio.NewWriter(os.Stdout))
		errCheck(encoder.Encode(records))
	}
}
