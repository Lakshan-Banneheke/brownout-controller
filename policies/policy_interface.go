package policies

type IPolicy interface {
	ExecuteForCluster()
	ExecuteForNode(nodeName string)
	sortPodsCluster() []string
	sortPodsNode(nodeName string) []string
}
