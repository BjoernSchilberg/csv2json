# Tools to convert CSV and JSON

- [Tools to convert CSV and JSON](#Tools-to-convert-CSV-and-JSON)
  - [csv2json](#csv2json)
    - [Get it](#Get-it)
    - [Use it](#Use-it)
    - [Format the output](#Format-the-output)
    - [Testing with real data](#Testing-with-real-data)
  - [json2csv](#json2csv)
    - [Get it](#Get-it-1)
    - [Use it](#Use-it-1)
  - [csv2xlsx](#csv2xlsx)
    - [Get it](#Get-it-2)
    - [Use it](#Use-it-2)
  - [License](#License)


## csv2json

- Reads CSV records from standard in.
- Outputs a JSON array of map entries to standard out.
- Assumes the first CSV record is the header.

### Get it

```shell
 go get -u bitbucket.org/BjoernSchilberg/csv2json/cmd/csv2json
```

Place the resulting `csv2json` in your `$PATH`.

### Use it

```shell
csv2json < file.csv > file.json
```

### Format the output

```shell
cat file.csv | ./csv2json -p
```

### Testing with real data

Real data for testing is for e.g. available from the USGS
[Earthquake Hazards Program](http://earthquake.usgs.gov/earthquakes/) as
[CSV](http://earthquake.usgs.gov/earthquakes/feed/v1.0/csv.php).

```shell
curl http://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_month.csv | csv2json > all_month.json
```

## json2csv

- Reads JSON array structure from standard in or optional an JSON object whichs holds the JSON array structure
- Outputs a CSV to standard out.

### Get it

```shell
go get -u bitbucket.org/BjoernSchilberg/csv2json/cmd/json2csv
```

Place the resulting `json2csv` in your `$PATH`.

### Use it

```shell
json2csv -h
```

```shell
json2csv < file.json > file.csv
```

```shell
json2csv -array=false --object=observations < file.json > file.csv
```

## csv2xlsx

- Reads CSV records from standard in.
- Outputs a xlsx to standard out.
- Assumes the first CSV record is the header.

### Get it

```shell
go get -u bitbucket.org/BjoernSchilberg/csv2json/cmd/csv2xlsx
```

Place the resulting `csv2xlsx` in your `$PATH`.

### Use it

```shell
csv2xlsx -h
```

```shell
csv2xlsx < file.csv > file.xlsx
```

```shell
csv2xlsx -sheet=observations < file.csv > file.xlsx
```

## License

This is Free Software under the terms of the MIT license.
See [LICENSE](LICENSE) file for details.
(c) 2016-2019 by Bjoern Schilberg
