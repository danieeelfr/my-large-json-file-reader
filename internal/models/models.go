package models

type Port struct {
	Name        interface{} `json:"name"`
	City        interface{} `json:"city"`
	Country     interface{} `json:"country"`
	Alias       interface{} `json:"alias"`
	Regions     interface{} `json:"regions"`
	Coordinates interface{} `json:"coordinates"`
	Provice     interface{} `json:"province"`
	Timezone    interface{} `json:"timezone"`
	Unlocs      interface{} `json:"unlocs"`
	Code        interface{} `json:"code"`
}
