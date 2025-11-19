package models

type ClientIP struct {
	IP       string  `json:"ip"`
	Country  string  `json:"country"`
	City     string  `json:"city"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Timezone string  `json:"timezone"`
	ISP      string  `json:"isp"`
}
