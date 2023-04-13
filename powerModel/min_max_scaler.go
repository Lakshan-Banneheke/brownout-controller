package powerModel

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

type MinMaxScaler struct {
	mins []float64
	maxs []float64
}

var scaler *MinMaxScaler
var once sync.Once

func GetScaler() *MinMaxScaler {
	once.Do(func() {
		scaler = &MinMaxScaler{}
		data := getDataFromFile("./powerModel/data/final-test-data-v1.csv")
		scaler.Fit(data)
	})
	return scaler
}

func (scaler *MinMaxScaler) Fit(data [][]float64) {
	numRows := len(data)
	numFeatures := len(data[0])

	for i := 0; i < numFeatures; i++ {
		scaler.mins = append(scaler.mins, data[0][i])
		scaler.maxs = append(scaler.maxs, data[0][i])
	}

	for i := 1; i < numRows; i++ {
		for j := 0; j < numFeatures; j++ {
			if data[i][j] < scaler.mins[j] {
				scaler.mins[j] = data[i][j]
			}
			if data[i][j] > scaler.maxs[j] {
				scaler.maxs[j] = data[i][j]
			}
		}
	}
}

func (scaler *MinMaxScaler) Transform(data []float64) []float64 {
	numFeatures := len(data)

	normalizedData := make([]float64, 6)

	for i := 0; i < numFeatures; i++ {
		normalizedData[i] = (data[i] - scaler.mins[i]) / (scaler.maxs[i] - scaler.mins[i])
	}

	return normalizedData
}

func getDataFromFile(filepath string) [][]float64 {
	// Open the CSV file
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

	var data [][]float64

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the CSV data into a slice of slices
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		floatRow := make([]float64, len(row))

		for i := range row {
			floatValue, err := strconv.ParseFloat(row[i], 64)
			if err != nil {
				continue
			}
			floatRow[i] = floatValue
		}

		data = append(data, floatRow)
	}

	return data
}
