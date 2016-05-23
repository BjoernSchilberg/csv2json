#csv2json - A simple tool to convert csv to json

* Reads CSV records from standard in.
* Outputs a JSON array of map entries to standard out. 
* Assumes the first CSV record is the header.

## Build

    $ go get -u bitbucket.org/BjoernSchilberg/csv2json 

Place the resulting `csv2json` binary into your PATH.

## Usage

    $ cat file.csv | csv2json

Or

    $ csv2json  < files.csv

Format the output:

    $ cat file.csv | ./csv2json | python -mjson.tool

## License
This is Free Software under the terms of the MIT license.
See [LICENSE](LICENSE) file for details.  
(c) 2016 by Bjoern Schilberg