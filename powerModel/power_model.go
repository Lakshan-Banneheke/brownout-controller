package powerModel

import (
	"brownout-controller/kubernetesCluster"
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

// map to store version with relevant power consumption calculating function for pods
var powerConsumptionPodsMap = map[string]func(*PowerModel, []string, string) float64{
	"v1": (*PowerModel).getPowerConsumptionPodsV1,
	"v2": (*PowerModel).getPowerConsumptionPodsV2,
	"v3": (*PowerModel).getPowerConsumptionPodsV3,
	"v4": (*PowerModel).getPowerConsumptionPodsV4,
}

func GetPowerModel() *PowerModel {
	oncePM.Do(func() {
		// initialize power model for the first time
		model = &PowerModel{}
		model.powerModelVersion = "v1"
		model.setCoefficients()
	})
	return model
}

func (model *PowerModel) GetPowerConsumptionNodes(nodeNames []string, namespace string) float64 {
	clientset, _ := kubernetesCluster.GetClientSets()                             // retrieve client set and metrics client
	podNames := kubernetesCluster.GetPodsInNodes(nodeNames, clientset, namespace) // retrieve all the pod names of the given nodes

	// call the pod power consumption calculating function for the relevant version
	return model.GetPowerConsumptionPods(podNames, namespace)
}

func (model *PowerModel) GetPowerConsumptionPods(podNames []string, namespace string) float64 {
	// call the pod power consumption calculating function for the relevant version
	return powerConsumptionPodsMap[model.powerModelVersion](model, podNames, namespace)
}

func (model *PowerModel) setCoefficients() {
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

func (model *PowerModel) calculatePower(params []float64) float64 {
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

func (model *PowerModel) getPowerConsumptionPodsV1(podNames []string, namespace string) float64 {

	/*
		Not yet implemented
		TODO: implement for V1 - add master memory and cpu
	*/
	clientset, metricsClient := kubernetesCluster.GetClientSets()                               // retrieve client set and metrics client
	podsCPUUsageSum := kubernetesCluster.GetPodsCPUUsageSum(metricsClient, podNames, namespace) // get the sum of CPU usage of the mentioned pods
	podsMemUsageSum := kubernetesCluster.GetPodsMemUsageSum(metricsClient, podNames, namespace) // get the sum of Memory usage of the mentioned pods
	workerNodeCount := float64(kubernetesCluster.GetWorkerNodeCount(clientset))                 // get the number of worker nodes
	podCount := float64(len(podNames))                                                          // calculate the pod count

	params := []float64{workerNodeCount, podCount, podsCPUUsageSum, podsMemUsageSum} //generate the input parameter list for calculating power

	power := model.calculatePower(params) // calculate the power using the model
	return power
}

func (model *PowerModel) getPowerConsumptionPodsV2(podNames []string, namespace string) float64 {
	clientset, metricsClient := kubernetesCluster.GetClientSets()                               // retrieve client set and metrics client
	podsCPUUsageSum := kubernetesCluster.GetPodsCPUUsageSum(metricsClient, podNames, namespace) // get the sum of CPU usage of the mentioned pods
	podsMemUsageSum := kubernetesCluster.GetPodsMemUsageSum(metricsClient, podNames, namespace) // get the sum of Memory usage of the mentioned pods
	workerNodeCount := float64(kubernetesCluster.GetWorkerNodeCount(clientset))                 // get the number of worker nodes
	podCount := float64(len(podNames))                                                          // calculate the pod count

	params := []float64{workerNodeCount, podCount, podsCPUUsageSum, podsMemUsageSum} //generate the input parameter list for calculating power

	power := model.calculatePower(params) // calculate the power using the model
	return power
}

func (model *PowerModel) getPowerConsumptionPodsV3(podNames []string, namespace string) float64 {
	clientset, metricsClient := kubernetesCluster.GetClientSets()                               // retrieve client set and metrics client
	podsCPUUsageSum := kubernetesCluster.GetPodsCPUUsageSum(metricsClient, podNames, namespace) // get the sum of CPU usage of the mentioned pods
	workerNodeCount := float64(kubernetesCluster.GetWorkerNodeCount(clientset))                 // get the number of worker nodes
	podCount := float64(len(podNames))                                                          // calculate the pod count

	params := []float64{workerNodeCount, podCount, podsCPUUsageSum} //generate the input parameter list for calculating power

	power := model.calculatePower(params) // calculate the power using the model
	return power
}

func (model *PowerModel) getPowerConsumptionPodsV4(podNames []string, namespace string) float64 {
	clientset, metricsClient := kubernetesCluster.GetClientSets()                               // retrieve client set and metrics client
	podsCPUUsageSum := kubernetesCluster.GetPodsCPUUsageSum(metricsClient, podNames, namespace) // get the sum of CPU usage of the mentioned pods
	workerNodeCount := float64(kubernetesCluster.GetWorkerNodeCount(clientset))                 // get the number of worker nodes

	params := []float64{workerNodeCount, podsCPUUsageSum} //generate the input parameter list for calculating power

	power := model.calculatePower(params) // calculate the power using the model
	return power
}
