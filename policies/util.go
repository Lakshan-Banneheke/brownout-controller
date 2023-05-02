package policies

import "fmt"

func GetSelectedPolicy(policyName string) IPolicy {
	switch policyName {
	case "NIMSP":
		return NIMSP{}
	case "NISP":
		return NISP{}
	case "LUCF":
		return LUCF{}
	case "HUCF":
		return HUCF{}
	case "RCSP":
		return RCSP{}
	default:
		panic(fmt.Sprintf("Error: Policy %s not available", policyName))
	}
}

func GetSelectedPodPolicy(policyName string) IPolicyPods {
	switch policyName {
	case "LUCF":
		return LUCF{}
	case "HUCF":
		return HUCF{}
	case "RCSP":
		return RCSP{}
	default:
		panic(fmt.Sprintf("Error: Policy %s not available", policyName))
	}
}
