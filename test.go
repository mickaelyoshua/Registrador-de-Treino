package main

import (
	"fmt"
	"log"
)

const (
	UnitCelsius    = "C"
	UnitFahrenheit = "F"
)

func convertFtoC(fahrenheit float64) float64 {
	const factor = 5.0 / 9.0
	return (fahrenheit - 32) * factor
}

func convertCtoF(celsius float64) float64 {
	const factor = 9.0 / 5.0
	return (celsius * factor) + 32
}

// function that takes either a celcius or fahrenheit value and returns the converted value
func convertTemperature(value float64, unit string) (float64, error) {
	switch unit {
	case UnitCelsius:
		return convertCtoF(value), nil
	case UnitFahrenheit:
		return convertFtoC(value), nil
	default:
		return 0, fmt.Errorf("invalid unit %s", unit)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func main() {
	// test the convertTemperature function
	celsius := 37.0
	fahrenheit, err := convertTemperature(celsius, "C")
	handleError(err)

	fmt.Printf("Temperature in Celsius: %.2f\n", celsius)
	fmt.Printf("Temperature in Fahrenheit: %.2f\n", fahrenheit)
	// test the convertTemperature function
	fahrenheit = 98.6
	celsius, err = convertTemperature(fahrenheit, "S")
	handleError(err)

	fmt.Printf("Temperature in Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Temperature in Celsius: %.2f\n", celsius)
}