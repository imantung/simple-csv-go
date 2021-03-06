package main

import (
	"flag"
	"log"
	"simple-csv-go/csv2"
)

var (
	path       = "example.csv"
	separator  = ";"
	withHeader = true
)

func main() {

	flag.Parse()

	// open: doublequote automatically remove
	reader := csv2.Reader{}
	err := reader.Open(path, separator, withHeader)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	// validate
	expectedColumns := []string{"column1", "column2"}
	if ok, err := reader.Validate(expectedColumns); !ok {
		log.Fatalf("Invalid csv: %s", err.Error())
	}

	// loop the rows
	for reader.Next() {
		row, err := reader.ReadAsRow()
		if err != nil {
			log.Printf("This row is ignored due to error: %s\n", err.Error())
		}

		rowNumber := row.RowNumber
		column1, _ := row.GetByName("column1")
		column2, _ := row.GetByName("column2")

		log.Printf("Line:%d 'name':%s 'age':%s", rowNumber, column1, column2)
	}

	log.Println("Done!")
}
