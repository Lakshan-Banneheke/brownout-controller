package powerModel

import (
	"fmt"
)

func CalculatePower(params []float64) {

	scaler := GetScaler()
	normalizedParams := scaler.Transform(params)
	for _, row := range normalizedParams {
		fmt.Println(row)
	}
}
