package util

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ExtractDataFromCSV(filepath string) [][]string {
	// open the CSV file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
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
