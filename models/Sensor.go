package models

type Sensor struct {
	Id                 string `json:"id"`
	Name               string
	MinimumTemperature int8
	MaximumTemperature int8
	Temperature        int8 `json:"temperature"`
}
