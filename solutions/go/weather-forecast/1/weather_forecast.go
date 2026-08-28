//Package weather provides tools to forcast the weather, based on the current location.
package weather

var (
	//CurrentCondition is a string representing the current weather conditions.
	CurrentCondition string
	//CurrentLocation is a string representing the Location.
	CurrentLocation  string
)
//Forecast takes in a city and the local weather conditions as strigs 
//and reformats it into a string focasting the weather at set location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
