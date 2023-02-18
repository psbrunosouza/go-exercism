// Package weather returns information about current wheter condition.
package weather

// CurrentCondition is the stored condition about the wheter.
var CurrentCondition string

// CurrentLocation is the localization in the city.
var CurrentLocation string

// Forecast receive two params city and condition and retuns how is the current wheter in a determined location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
