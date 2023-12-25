// Package weather gives
// information about the weather.
package weather

// CurrentCondition gives the current weather.
var CurrentCondition string
// CurrentLocation gives the current location.
var CurrentLocation string

// Forecast returns a location with the current weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
