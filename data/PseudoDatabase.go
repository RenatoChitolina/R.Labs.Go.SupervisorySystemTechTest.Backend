package data

import "server/models"

func SensorSetups() map[string]models.Sensor {
	var data = map[string]models.Sensor{
		"1": {
			Id:                 "1",
			Name:               "Pilsner",
			MinimumTemperature: 4,
			MaximumTemperature: 6,
		},
		"2": {
			Id:                 "2",
			Name:               "IPA",
			MinimumTemperature: 5,
			MaximumTemperature: 6,
		},
		"3": {
			Id:                 "3",
			Name:               "Lager",
			MinimumTemperature: 4,
			MaximumTemperature: 7,
		},
		"4": {
			Id:                 "4",
			Name:               "Stout",
			MinimumTemperature: 6,
			MaximumTemperature: 8,
		},
		"5": {
			Id:                 "5",
			Name:               "Wheat beer",
			MinimumTemperature: 3,
			MaximumTemperature: 5,
		},
		"6": {
			Id:                 "6",
			Name:               "Pale Ale",
			MinimumTemperature: 4,
			MaximumTemperature: 6,
		},
		"NonNumericKey": {
			Id:                 "NonNumericKey",
			Name:               "Export",
			MinimumTemperature: 2,
			MaximumTemperature: 2,
		}}

	return data
}
