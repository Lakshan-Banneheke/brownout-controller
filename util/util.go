package util

import (
	"brownout-controller/policies"
	"fmt"
)

// SliceDifference returns the difference between two string slices.
// The difference here means that the values which are in slice1 but not present in slice2.
func SliceDifference(slice1, slice2 []string) []string {
	var difference []string
	m := make(map[string]bool)

	for _, val := range slice2 {
		m[val] = true
	}

	for _, val := range slice1 {
		if _, ok := m[val]; !ok {
			difference = append(difference, val)
		}
	}
	return difference
}

// AddDeployments functions appends values from the map deployments1 to the map deployments2 and returns deployments2
func AddDeployments(deployments1 map[string]int32, deployments2 map[string]int32) map[string]int32 {
	for key, value1 := range deployments1 {
		if value2, exists := deployments2[key]; exists {
			deployments2[key] = value2 + value1
		} else {
			deployments2[key] = value1
		}
	}
	return deployments2
}

func GetSelectedPolicy(policyName string) policies.IPolicy {
	switch policyName {
	case "NISP":
		return policies.NISP{}
	case "LUCF":
		return policies.LUCF{}
	case "HUCF":
		return policies.HUCF{}
	case "RCSP":
		return policies.RCSP{}
	default:
		panic(fmt.Sprintf("Error: Policy %s not available", policyName))
	}
}

func GetSelectedPodPolicy(policyName string) policies.IPolicyPods {
	switch policyName {
	case "LUCF":
		return policies.LUCF{}
	case "HUCF":
		return policies.HUCF{}
	case "RCSP":
		return policies.RCSP{}
	default:
		panic(fmt.Sprintf("Error: Policy %s not available", policyName))
	}
}
