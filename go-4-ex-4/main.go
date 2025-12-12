package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(celcius float64) float64 {
	return celcius*9/5 + 32
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	f := convertCelsiusToFahrenheit(24)
	fmt.Println(f, "°F")
	// TODO: call the function convertFahrenheitToCelsius
	c := convertFahrenheitToCelsius(108)
	fmt.Println(c, "°C")
}
