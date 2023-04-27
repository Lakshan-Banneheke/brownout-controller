package policies

// IPolicy for all pod selection policies
type IPolicy interface {
}

// IPolicyPods for all pod selection policies that are pod-wise (lucf, lru, rcsp)
type IPolicyPods interface {
	IPolicy
	ExecuteForCluster(upperThresholdPower float64) map[string]int32
	ExecuteForNode(nodeName string, upperThresholdPower float64) map[string]int32
	sortPodsCluster() []string
	sortPodsNode(nodeName string) []string
}

// IPolicyNodes for all pod selection policies that are node-wise (nisp)
type IPolicyNodes interface {
	IPolicy
	ExecuteForCluster(upperThresholdPower float64) (map[string]int32, []string)
	deactivateNodes(nodeList []string) []string
}
