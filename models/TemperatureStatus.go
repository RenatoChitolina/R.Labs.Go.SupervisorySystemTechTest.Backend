package models

type TemperatureStatus string

var TemperatureStatuses = struct {
	TooLow  string
	TooHigh string
	AllGood string
}{
	TooLow:  "too low",
	TooHigh: "too high",
	AllGood: "all good",
}
