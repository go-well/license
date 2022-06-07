package activation_code

import (
	"time"
)

type Licence struct {
	Product   string
	User      string
	UUID      string `json:"uuid"`
	SN        string `json:"sn"`
	CpuId     string `json:"cpuId"`
	MAC       string `json:"mac"`
	ExpireAt  time.Time
	Signature string
}
