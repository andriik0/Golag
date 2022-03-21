// Package contains string variables CurrentCondition and CurrentLocation and Forecast function
package weather
// Exported variable store current weather condition as string that can be read or write by user
var CurrentCondition string
// Exported variable store current forecast location as string that can be read or write by user
var CurrentLocation string

// Forecast function return forecast for defined city and condition
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
