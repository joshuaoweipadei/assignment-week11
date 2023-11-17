package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiKey = "6dbc65c242ded0d9ca9dd9e3a7035d22" // OpenWeatherMap API key

type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

// Define a custom struct representing an object
type WeatherObj struct {
	Description string
	Temp        int
}

func torontoWeather() WeatherObj {
	city := "Toronto"
	weatherURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	response, err := http.Get(weatherURL)
	if err != nil {
		fmt.Printf("Error retriving data: %s\n", err)
		return WeatherObj{}
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return WeatherObj{}
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
		return WeatherObj{}
	}

	return WeatherObj{Description: weatherData.Weather[0].Description, Temp: int(weatherData.Main.Temp - 273.15)} // Convert Kelvin to Celsius
}

func torontoTimeHandler(w http.ResponseWriter, r *http.Request) {
	torontoWeatherObj := torontoWeather()

	// fmt.Fprintf(w, "Toronto time is", torontoWeatherObj.Description, torontoWeatherObj.Temp)
	resp := map[string]interface{}{"Temperature": torontoWeatherObj.Temp, "Description": torontoWeatherObj.Description}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/toronto-weather", torontoTimeHandler)
	fmt.Println("Server is listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
