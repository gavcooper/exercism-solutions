// Package weather contains functions to help forecast the weather of a particular city in Goblinocus.
package weather

// CurrentCondition weather condition.
var CurrentCondition string

// CurrentLocation is a string with the current location.
var CurrentLocation string

// Forecast takes a city and condition and returns a string describing the current weather in a particular city. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}