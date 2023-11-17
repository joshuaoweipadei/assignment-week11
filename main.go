package main

import (
	"fmt"
	"net/http"
	"log"
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
	Temp int
}

func main()  {
	fmt.Println("Server is listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}