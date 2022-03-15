package models

type Sensor struct {
	Id                 string
	Name               string
	MinimumTemperature int8
	MaximumTemperature int8
	Temperature        int8 `json:"temperature"`
	Status             TemperatureStatus
}

func (sensor *Sensor) CheckStatus() {
	sensor.Status = TemperatureStatus(TemperatureStatuses.AllGood)
}
