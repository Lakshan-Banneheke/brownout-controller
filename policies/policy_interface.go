package policies

// IPolicy for all pod selection policies
type IPolicy interface {
	ExecuteForCluster()
}

// IPolicyPods for all pod selection policies that are pod-wise (lucf, lru, rcsp)
type IPolicyPods interface {
	IPolicy
	ExecuteForNode(nodeName string)
	sortPodsCluster() []string
	sortPodsNode(nodeName string) []string
}

// IPolicyNodes for all pod selection policies that are node-wise (nisp)
type IPolicyNodes interface {
	IPolicy
	deactivateNodes(nodeList []string)
}
