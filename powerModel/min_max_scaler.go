package powerModel

import (
	"brownout-controller/powerModel/util"
	"log"
	"strconv"
	"sync"
)

type MinMaxScaler struct {
	mins []float64
	maxs []float64
}

var scaler *MinMaxScaler
var once sync.Once

func GetScaler(version string) *MinMaxScaler {
	once.Do(func() {
		// initialize scaler for the first time
		scaler = &MinMaxScaler{}
		data := getDataFromFile("./powerModel/data/final-test-data-" + version + ".csv")
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

	normalizedData := make([]float64, numFeatures)

	for i := 0; i < numFeatures; i++ {
		normalizedData[i] = (data[i] - scaler.mins[i]) / (scaler.maxs[i] - scaler.mins[i])
	}

	return normalizedData
}

func getDataFromFile(filepath string) [][]float64 {
	// extract data needed to fit the scaler from the csv file
	rows := util.ExtractDataFromCSV(filepath)

	var data [][]float64

	// convert the data into floats and populate the slice
	for j, row := range rows {
		if j == 0 {
			// skip the header row
			continue
		}

		floatRow := make([]float64, len(row))

		for i := range row {
			floatValue, err := strconv.ParseFloat(row[i], 64)
			if err != nil {
				log.Fatal(err)
			}
			floatRow[i] = floatValue
		}

		data = append(data, floatRow)
	}

	return data
}
