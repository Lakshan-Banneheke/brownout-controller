package util

import (
	"embed"
	"encoding/csv"
	"io"
	"io/fs"
	"log"
)

//go:embed data/*
var fileContent embed.FS

// ExtractDataFromCSV : function to extract the rows of csv file
func ExtractDataFromCSV(filepath string) [][]string {

	// open the CSV file
	file, err := fileContent.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file fs.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// create a new CSV reader
	reader := csv.NewReader(file)

	var rows [][]string

	// read and store the CSV data into a slice of slices
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		rows = append(rows, row)
	}

	return rows
}
