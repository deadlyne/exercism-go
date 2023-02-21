// Package weather got information about weather.
package weather

// CurrentCondition got information about Condition.
var CurrentCondition string

// CurrentLocation got information about Location.
var CurrentLocation string

// Forecast return City and condition in the City.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
