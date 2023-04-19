package powerModel

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"sync"
)

type PowerModel struct {
	coefficients []float64
}

var model *PowerModel
var oncePM sync.Once

// GetPowerModel : function to retrieve the power model
func GetPowerModel() *PowerModel {

	oncePM.Do(func() {
		// initialize power model
		model = &PowerModel{}
		model.setCoefficients()
	})
	return model
}

// GetPowerConsumptionNodes : function to compute power consumption when a set of nodes given
func (model *PowerModel) GetPowerConsumptionNodes(nodeNames []string) float64 {

	podNames := kubernetesCluster.GetPodsInNodes(nodeNames, constants.NAMESPACE) // retrieve all the pod names of the given nodes
	workerNodeCount := float64(len(nodeNames))                                   // get the number of worker nodes
	return getPower(podNames, workerNodeCount, model)
}

// GetPowerConsumptionPods : function to compute power consumption when a set of pods given
func (model *PowerModel) GetPowerConsumptionPods(podNames []string) float64 {

	workerNodeCount := float64(kubernetesCluster.GetWorkerNodeCount()) // get the number of worker nodes
	return getPower(podNames, workerNodeCount, model)
}

// function to set coefficients of the power model
func (model *PowerModel) setCoefficients() {

	model.coefficients = []float64{constants.C1, constants.C2, constants.C3, constants.C4}
}

func getPower(podNames []string, workerNodeCount float64, model *PowerModel) float64 {
	// retrieve input parameters needed by the power model
	podsCPUUsageSum := kubernetesCluster.GetPodsCPUUsageSum(podNames, constants.NAMESPACE) // get the sum of CPU usage of the mentioned pods
	podCount := float64(len(podNames))                                                     // calculate the pod count

	//generate the input parameter list for calculating power
	params := []float64{workerNodeCount, podCount, podsCPUUsageSum}

	// calculate the power using the model
	power := model.calculatePower(params)
	return power
}

// function to calculate the power from linear regression model
func (model *PowerModel) calculatePower(params []float64) float64 {

	scaler := GetScaler()                                        // get the min max scaler
	normalizedParams := scaler.Transform(params)                 // normalize the input parameters
	normalizedParams = append([]float64{1}, normalizedParams...) // append 1 to the front to facilitate the bias term

	// predict power using the coefficients of the linear regression model
	power := 0.0
	for i, coefficient := range model.coefficients {
		power += coefficient * normalizedParams[i]
	}

	return power
}
