package main

import (
	"fmt"
	"os"
)

func convertFtoC(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
func convertCtoF(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// function that takes either a celcius or fahrenheit value and returns the converted value
func convertTemperature(value float64, unit string) (float64, error) {
	switch unit {
	case "C":
		return convertCtoF(value), nil
	case "F":
		return convertFtoC(value), nil
	default:
		return 0, fmt.Errorf("invalid unit %s", unit)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
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