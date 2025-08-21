//Package weather provites tools to display the current weather at a given location.
package weather

// CurrentCondition representes the current weather conditions.
var CurrentCondition string
// CurrentLocation represents the current location.
var CurrentLocation string


// Forecast return a string value equalt to the given CurrentLocation and CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
