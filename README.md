#csv2json - A simple tool to convert csv to json

* Reads CSV records from standard in.
* Outputs a JSON array of map entries to standard out. 
* Assumes the first CSV record is the header.

## Build

    $ go get -u bitbucket.org/BjoernSchilberg/csv2json/cmd/csv2json

or if you want to go the other way:

    $ go get -u bitbucket.org/BjoernSchilberg/csv2json/cmd/json2csv

Place the resulting `csv2json` and `json2csv` binaries into your PATH.

## Usage

    $ csv2json < file.csv > file.json
    $ json2csv < file.json > file.csv

Format the output:

    $ cat file.csv | ./csv2json | python -mjson.tool

## License
This is Free Software under the terms of the MIT license.
See [LICENSE](LICENSE) file for details.  
(c) 2016 by Bjoern Schilberg
