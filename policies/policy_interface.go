package policies

// IPolicy for all pod selection policies
type IPolicy interface {
	ExecuteForCluster(upperThresholdPower float64) map[string]int32
}

// IPolicyPods for all pod selection policies that are pod-wise (lucf, lru, rcsp)
type IPolicyPods interface {
	IPolicy
	ExecuteForNode(nodeName string, upperThresholdPower float64) map[string]int32
	sortPodsCluster() []string
	sortPodsNode(nodeName string) []string
}

// IPolicyNodes for all pod selection policies that are node-wise (nisp)
type IPolicyNodes interface {
	IPolicy
	deactivateNodes(nodeList []string) map[string]int32
}
