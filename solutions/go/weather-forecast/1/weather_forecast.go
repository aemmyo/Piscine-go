// Package weather provide tools for checking the current weather condition.
package weather

// CurrentCondition stores the current weather condition.  
var CurrentCondition string
    
// CurrentLocation stores the current location of that weather.
var CurrentLocation  string 

// Forecast returns the string value of the place with that weather condition. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
