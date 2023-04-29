package brownout

var batteryPercentage int

// User can use this API endpoint to periodically send the battery percentage to the brownout controller
func setBatteryPercentage(y int) {
	batteryPercentage = y
}

func getBatteryPercentage() int {
	return batteryPercentage
}
