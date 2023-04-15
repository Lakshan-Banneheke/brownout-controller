package main

import (
	"brownout-controller/powerModel"
	"fmt"
)

func main() {
	// get the power model
	pm := powerModel.GetPowerModel()

	fmt.Println(pm.GetPowerConsumptionPods([]string{"harry", "potter"}, "default"))

}
