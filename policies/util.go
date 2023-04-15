package policies

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
