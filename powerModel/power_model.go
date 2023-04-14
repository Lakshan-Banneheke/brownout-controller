package powerModel

import (
	"brownout-controller/powerModel/util"
	"log"
	"strconv"
)

var coefficients []float64
var powerModelVersion string

func SetCoefficients(version string) {
	// initialize the version
	powerModelVersion = version

	// extract coefficients from csv file
	rows := util.ExtractDataFromCSV("./powerModel/data/coefficients/analytical-model-lr-" + powerModelVersion + "-coefficients.csv")

	// convert the coefficients to floats and populate the slice
	for _, row := range rows {
		coefficient, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		coefficients = append(coefficients, coefficient)
	}
}

func CalculatePower(params []float64) float64 {
	scaler := GetScaler(powerModelVersion)                       // get the min max scaler
	normalizedParams := scaler.Transform(params)                 // normalize the input parameters
	normalizedParams = append([]float64{1}, normalizedParams...) // append 1 to the front to facilitate the bias term

	// predict power
	power := 0.0
	for i, coefficient := range coefficients {
		power += coefficient * normalizedParams[i]
	}

	return power
}
