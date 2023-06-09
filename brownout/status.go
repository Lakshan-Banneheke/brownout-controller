package brownout

var brownoutActive bool

func SetBrownoutActive(y bool) {
	brownoutActive = y
}

func GetBrownoutActive() bool {
	return brownoutActive
}
