package variables

var (
	POLICY                    = "NISP" // Options: NISP, LUCF, HUCF, RCSP
	BATTERY_UPPER_THRESHOLD   = 80
	BATTERY_LOWER_THRESHOLD   = 50
	SLA_VIOLATION_LATENCY     = "0.25" //seconds
	SLA_INTERVAL              = "1m"
	ACCEPTED_SUCCESS_RATE     = 0.65
	ACCEPTED_MIN_SUCCESS_RATE = 0.50
)
