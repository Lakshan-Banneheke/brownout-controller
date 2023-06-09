package kubernetesCluster

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

// function to retrieve the CPU Usage metrics from a given set of metrics
func extractCPUMetrics(podMetricsItems []v1beta1.PodMetrics, err error) (map[string]int, []string) {

	if err != nil {
		fmt.Println(err.Error())
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
				fmt.Println(err.Error())
			}
			podCPU += cpuUsageInt
		}

		podsCPUUsage[podMetric.ObjectMeta.Name] = podCPU
		podNames = append(podNames, podMetric.ObjectMeta.Name)
	}

	return podsCPUUsage, podNames
}

// function returns a list of node names in sorted order of increasing cpu usage
func sortNodesUsageAscending(nodesCPUUsage map[string]int, nodeNames []string) []string {

	sort.SliceStable(nodeNames, func(i, j int) bool {
		return nodesCPUUsage[nodeNames[i]] < nodesCPUUsage[nodeNames[j]]
	})

	return nodeNames
}

// function returns a list of node names in sorted order of increasing cpu usage
func sortPodsUsageAscending(podsCPUUsage map[string]int, podNames []string) []string {

	sort.SliceStable(podNames, func(i, j int) bool {
		return podsCPUUsage[podNames[i]] < podsCPUUsage[podNames[j]]
	})

	return podNames
}

// function returns a list of node names in sorted order of decreasing cpu usage
func sortPodsUsageDescending(podsCPUUsage map[string]int, podNames []string) []string {

	sort.SliceStable(podNames, func(i, j int) bool {
		return podsCPUUsage[podNames[i]] > podsCPUUsage[podNames[j]]
	})

	return podNames
}

// function returns a list of pod names in a random order
func sortPodsRandomly(podNames []string) []string {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := len(podNames) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		podNames[i], podNames[j] = podNames[j], podNames[i]
	}

	return podNames
}
