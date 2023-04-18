package powerModel

import (
	"brownout-controller/constants"
	"sync"
)

type MinMaxScaler struct {
	mins []float64
	maxs []float64
}

var scaler *MinMaxScaler
var onceMMS sync.Once

// GetScaler : function to retrieve the Min Max Scaler
func GetScaler() *MinMaxScaler {

	onceMMS.Do(func() {
		// initialize scaler
		scaler = &MinMaxScaler{}
		fitData := [][]float64{{constants.MAX0, constants.MAX1, constants.MAX2}, {constants.MIN0, constants.MIN1, constants.MIN2}}
		scaler.fit(fitData)
	})
	return scaler
}

// Transform : function to normalize the given data
func (scaler *MinMaxScaler) Transform(data []float64) []float64 {

	numFeatures := len(data)

	normalizedData := make([]float64, numFeatures)

	for i := 0; i < numFeatures; i++ {
		normalizedData[i] = (data[i] - scaler.mins[i]) / (scaler.maxs[i] - scaler.mins[i])
	}

	return normalizedData
}

// function to fit the scaler with the relevant dataset
func (scaler *MinMaxScaler) fit(data [][]float64) {

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
