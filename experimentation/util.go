package experimentation

import (
	"brownout-controller/policies"
	"fmt"
)

func average(listFloat []float64) float64 {
	sum := 0.0
	for _, x := range listFloat {
		sum += x
	}
	return sum / float64(len(listFloat))
}

func getSelectedPolicy(policyName string) policies.IPolicy {
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
