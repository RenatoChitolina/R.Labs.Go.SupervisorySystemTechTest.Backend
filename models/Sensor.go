package models

type Sensor struct {
	Id                 string            `json:"id"`
	Name               string            `json:"name"`
	MinimumTemperature int8              `json:"minimumTemperature"`
	MaximumTemperature int8              `json:"maximumTemperature"`
	Temperature        int8              `json:"temperature"`
	Status             TemperatureStatus `json:"status"`
}

func (sensor *Sensor) CheckStatus() {
	if sensor.Temperature < sensor.MinimumTemperature {
		sensor.Status = TemperatureStatus(TemperatureStatuses.TooLow)
	} else if sensor.Temperature > sensor.MaximumTemperature {
		sensor.Status = TemperatureStatus(TemperatureStatuses.TooHigh)
	} else if sensor.Temperature <= sensor.MaximumTemperature &&
		sensor.Temperature >= sensor.MinimumTemperature { // Although we don't need this long "if" here (a simple "else" would solve), I'm keeping the rule exactly as same as was in the frontend
		sensor.Status = TemperatureStatus(TemperatureStatuses.AllGood)
	}
}
