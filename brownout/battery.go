package brownout

import "log"

var batteryPercentage int

func SetBatteryPercentage(y int) {
	log.Printf("Battery percentage set to %v%%", y)
	batteryPercentage = y
}

func GetBatteryPercentage() int {
	return batteryPercentage
}
