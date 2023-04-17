package powerModel

import (
	"brownout-controller/constants"
	"brownout-controller/kubernetesCluster"
	"brownout-controller/powerModel/util"
	"log"
	"strconv"
)

type PowerModel struct {
	coefficients      []float64
	powerModelVersion string
}

var model *PowerModel

// map to store version with relevant power consumption calculating function for pods
var powerConsumptionPodsMap = map[string]func(*PowerModel, []string, string) float64{
	"v1": (*PowerModel).getPowerConsumptionPodsV1,
	"v2": (*PowerModel).getPowerConsumptionPodsV2,
	"v3": (*PowerModel).getPowerConsumptionPodsV3,
	"v4": (*PowerModel).getPowerConsumptionPodsV4,
}

// GetPowerModel : function to retrieve the power model
func GetPowerModel(version string) *PowerModel {

	if model == nil || model.powerModelVersion != version {
		// initialize power model
		model = &PowerModel{
			powerModelVersion: version,
		}
		model.setCoefficients()
	}
	return model
}

// GetPowerConsumptionNodes : function to compute power consumption when a set of nodes given
func (model *PowerModel) GetPowerConsumptionNodes(nodeNames []string) float64 {

	podNames := kubernetesCluster.GetPodsInNodes(nodeNames, constants.NAMESPACE) // retrieve all the pod names of the given nodes

	// call the pod power consumption calculating function for the relevant version
	return model.GetPowerConsumptionPods(podNames)
}

// GetPowerConsumptionPods : function to compute power consumption when a set of pods given
func (model *PowerModel) GetPowerConsumptionPods(podNames []string) float64 {

	// call the pod power consumption calculating function for the relevant version
	return powerConsumptionPodsMap[model.powerModelVersion](model, podNames, constants.NAMESPACE)
}

// function to set coefficients of the power model
func (model *PowerModel) setCoefficients() {

	// extract coefficients from csv file
	rows := util.ExtractDataFromCSV("data/coefficients/analytical-model-lr-" + model.powerModelVersion + "-coefficients.csv")

	// convert the coefficients to floats and populate the slice
	for _, row := range rows {
		coefficient, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		model.coefficients = append(model.coefficients, coefficient)
	}
}

// function to calculate the power from linear regression model
func (model *PowerModel) calculatePower(params []float64) float64 {

	scaler := GetScaler(model.powerModelVersion)                 // get the min max scaler
	normalizedParams := scaler.Transform(params)                 // normalize the input parameters
	normalizedParams = append([]float64{1}, normalizedParams...) // append 1 to the front to facilitate the bias term

	// predict power using the coefficients of the linear regression model
	power := 0.0
	for i, coefficient := range model.coefficients {
		power += coefficient * normalizedParams[i]
	}

	return power
}

// power calculation function - V1
func (model *PowerModel) getPowerConsumptionPodsV1(podNames []string, namespace string) float64 {

	// retrieve input parameters needed by the power model
	masterCPUUsage, masterMemUsage := kubernetesCluster.GetMasterNodeUsage()     // retrieve master node CPU and Memory usage
	podsCPUUsageSum := kubernetesCluster.GetPodsCPUUsageSum(podNames, namespace) // get the sum of CPU usage of the mentioned pods
	podsMemUsageSum := kubernetesCluster.GetPodsMemUsageSum(podNames, namespace) // get the sum of Memory usage of the mentioned pods
	workerNodeCount := float64(kubernetesCluster.GetWorkerNodeCount())           // get the number of worker nodes
	podCount := float64(len(podNames))                                           // calculate the pod count

	//generate the input parameter list for calculating power
	params := []float64{masterCPUUsage, masterMemUsage, workerNodeCount, podCount, podsCPUUsageSum, podsMemUsageSum}

	// calculate the power using the model
	power := model.calculatePower(params)
	return power
}

// power calculation function - V2
func (model *PowerModel) getPowerConsumptionPodsV2(podNames []string, namespace string) float64 {

	// retrieve input parameters needed by the power model
	podsCPUUsageSum := kubernetesCluster.GetPodsCPUUsageSum(podNames, namespace) // get the sum of CPU usage of the mentioned pods
	podsMemUsageSum := kubernetesCluster.GetPodsMemUsageSum(podNames, namespace) // get the sum of Memory usage of the mentioned pods
	workerNodeCount := float64(kubernetesCluster.GetWorkerNodeCount())           // get the number of worker nodes
	podCount := float64(len(podNames))                                           // calculate the pod count

	//generate the input parameter list for calculating power
	params := []float64{workerNodeCount, podCount, podsCPUUsageSum, podsMemUsageSum}

	// calculate the power using the model
	power := model.calculatePower(params)
	return power
}

// power calculation function - V3
func (model *PowerModel) getPowerConsumptionPodsV3(podNames []string, namespace string) float64 {

	// retrieve input parameters needed by the power model
	podsCPUUsageSum := kubernetesCluster.GetPodsCPUUsageSum(podNames, namespace) // get the sum of CPU usage of the mentioned pods
	workerNodeCount := float64(kubernetesCluster.GetWorkerNodeCount())           // get the number of worker nodes
	podCount := float64(len(podNames))                                           // calculate the pod count

	//generate the input parameter list for calculating power
	params := []float64{workerNodeCount, podCount, podsCPUUsageSum}
	log.Println(params)

	// calculate the power using the model
	power := model.calculatePower(params)
	return power
}

// power calculation function - V4
func (model *PowerModel) getPowerConsumptionPodsV4(podNames []string, namespace string) float64 {

	// retrieve input parameters needed by the power model
	podsCPUUsageSum := kubernetesCluster.GetPodsCPUUsageSum(podNames, namespace) // get the sum of CPU usage of the mentioned pods
	workerNodeCount := float64(kubernetesCluster.GetWorkerNodeCount())           // get the number of worker nodes

	//generate the input parameter list for calculating power
	params := []float64{workerNodeCount, podsCPUUsageSum}

	// calculate the power using the model
	power := model.calculatePower(params)
	return power
}
