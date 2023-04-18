package kubernetesCluster

import (
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"sort"
	"strconv"
)

// function to retrieve the CPU Usage metrics from a given set of metrics
func extractCPUMetrics(podMetricsItems []v1beta1.PodMetrics, err error) (map[string]int, []string) {

	if err != nil {
		panic(err.Error())
	}

	// make a map of pod Name and cpu usage
	podsCPUUsage := map[string]int{}
	var podNames []string

	for _, podMetric := range podMetricsItems {
		podCPU := 0
		for _, cont := range podMetric.Containers {
			contCPU := cont.Usage.Cpu().String()
			var contCPUTrimmed string
			// Removing the unit "n" from CPU Usage and converting to int for ease of sorting
			if contCPU != "0" {
				contCPUTrimmed = contCPU[:len(contCPU)-1]
			} else {
				contCPUTrimmed = contCPU
			}
			cpuUsageInt, err := strconv.Atoi(contCPUTrimmed)
			if err != nil {
				panic(err.Error())
			}
			podCPU += cpuUsageInt
		}

		podsCPUUsage[podMetric.ObjectMeta.Name] = podCPU
		podNames = append(podNames, podMetric.ObjectMeta.Name)
	}

	return podsCPUUsage, podNames
}

// function returns a list of node names in sorted order of increasing cpu usage
func sortNodesUsage(nodesCPUUsage map[string]int, nodeNames []string) []string {

	sort.SliceStable(nodeNames, func(i, j int) bool {
		return nodesCPUUsage[nodeNames[i]] < nodesCPUUsage[nodeNames[j]]
	})

	return nodeNames
}

// function returns a list of node names in sorted order of increasing cpu usage
func sortPodsUsage(podsCPUUsage map[string]int, podNames []string) []string {

	sort.SliceStable(podNames, func(i, j int) bool {
		return podsCPUUsage[podNames[i]] < podsCPUUsage[podNames[j]]
	})

	return podNames
}
