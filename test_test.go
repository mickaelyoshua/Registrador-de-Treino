package main

import (
	"testing"
)

func TestConvertFtoC(t *testing.T) {
	result := convertFtoC(32)
	if result != 0 {
		t.Errorf("Expected 0, got %f", result)
	}

	result = convertFtoC(212)
	if result != 100 {
		t.Errorf("Expected 100, got %f", result)
	}
}

func TestConvertCtoF(t *testing.T) {
	result := convertCtoF(0)
	if result != 32 {
		t.Errorf("Expected 32, got %f", result)
	}

	result = convertCtoF(100)
	if result != 212 {
		t.Errorf("Expected 212, got %f", result)
	}
}

func TestConvertTemperature(t *testing.T) {
	result, err := convertTemperature(0, UnitCelsius)
	if err != nil || result != 32 {
		t.Errorf("Expected 32, got %f, error: %v", result, err)
	}

	result, err = convertTemperature(32, UnitFahrenheit)
	if err != nil || result != 0 {
		t.Errorf("Expected 0, got %f, error: %v", result, err)
	}

	_, err = convertTemperature(100, "Invalid")
	if err == nil {
		t.Error("Expected error for invalid unit, got nil")
	}
}
