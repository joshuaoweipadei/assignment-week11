package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Declare and initialize the weatherURL variable
var weatherURL string

func TestTorontoWeather(t *testing.T) {
	// Mock the OpenWeatherMap API response
	mockResponse := `{"weather":[{"description":"Clear"}],"main":{"temp":280}}`

	// Create a test server to serve the mock response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer ts.Close()

	// Call the function to be tested
	result := torontoWeather()

	// Check if the result matches the expected values
	expected := WeatherObj{Description: "Clear", Temp: 6} // 280K to Celsius
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
