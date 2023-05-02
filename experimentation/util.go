package experimentation

func average(listFloat []float64) float64 {
	sum := 0.0
	for _, x := range listFloat {
		sum += x
	}
	return sum / float64(len(listFloat))
}
