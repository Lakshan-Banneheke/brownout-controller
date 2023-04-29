package brownout

import (
	"brownout-controller/policies"
	"fmt"
)

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
