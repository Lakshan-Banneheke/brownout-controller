package powerModel

import (
	"brownout-controller/powerModel/util"
	"log"
	"strconv"
	"sync"
)

type PowerModel struct {
	coefficients      []float64
	powerModelVersion string
}

var model *PowerModel
var oncePM sync.Once

func GetPowerModel(version string) *PowerModel {
	oncePM.Do(func() {
		// initialize power model for the first time
		model = &PowerModel{}
		model.setCoefficients(version)
	})
	return model
}

func (model *PowerModel) setCoefficients(version string) {
	// initialize the version
	model.powerModelVersion = version

	// extract coefficients from csv file
	rows := util.ExtractDataFromCSV("./powerModel/data/coefficients/analytical-model-lr-" + model.powerModelVersion + "-coefficients.csv")

	// convert the coefficients to floats and populate the slice
	for _, row := range rows {
		coefficient, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		model.coefficients = append(model.coefficients, coefficient)
	}
}

func (model *PowerModel) CalculatePower(params []float64) float64 {
	scaler := GetScaler(model.powerModelVersion)                 // get the min max scaler
	normalizedParams := scaler.Transform(params)                 // normalize the input parameters
	normalizedParams = append([]float64{1}, normalizedParams...) // append 1 to the front to facilitate the bias term

	// predict power
	power := 0.0
	for i, coefficient := range model.coefficients {
		power += coefficient * normalizedParams[i]
	}

	return power
}
