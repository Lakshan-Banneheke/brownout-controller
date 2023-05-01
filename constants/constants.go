package constants

const (
	POLICY                    = "NISP" // Options: NISP, LUCF, HUCF, RCSP
	NISP_PER_NODE_POLICY      = "LUCF"
	OPTIONAL                  = "optional"
	NAMESPACE                 = "default"
	BATTERY_UPPER_THRESHOLD   = 80
	BATTERY_LOWER_THRESHOLD   = 50
	SLA_VIOLATION_LATENCY     = "0.25" //seconds
	SLA_INTERVAL              = "1m"
	ACCEPTED_SUCCESS_RATE     = 0.65
	ACCEPTED_MIN_SUCCESS_RATE = 0.50
	K_VALUE                   = 0.0000 // Select a policy and put the respective K value
	K_LUCF                    = 0.9217
	K_HUCF                    = 0.9567
	K_RCSP                    = 0.9582
	K_NISP                    = 0.8199
	HOSTNAME                  = "agrimaster.com"
)
