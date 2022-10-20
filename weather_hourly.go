package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type weather interface {
	CurrentTemperature() float64
	CurrentWind() float64
	CurrentRain() float64
}

type weatherInfo struct {
	temperature float64
	wind        float64
	rain        float64
}

func CoallesceWeatherInfo(w weather) weatherInfo {
	var weatherData weatherInfo

	weatherData.temperature = w.CurrentTemperature()
	weatherData.wind = w.CurrentWind()
	weatherData.rain = w.CurrentRain()

	return weatherData
}

func (result *result) CurrentTemperature() float64 {
	return result.Hourly.Temperature[0]

}

func (result *result) CurrentWind() float64 {
	return result.Hourly.Windspeed[0]

}

func (result *result) CurrentRain() float64 {
	return result.Hourly.Rain[0]

}

func buildURL(base string, end string, coord coordinates) string {
	var fullURL string

	fullURL = base + "latitude=" + fmt.Sprintf("%v", coord.latitude) + "&longitude=" + fmt.Sprintf("%v", coord.longitude) + end

	return fullURL
}

type coordinates struct {
	longitude float64
	latitude  float64
}

type result struct {
	Hourly hourly `json:"hourly"`
}

type hourly struct {
	Temperature []float64 `json:"temperature_2m"`
	Windspeed   []float64 `json:"windspeed_10m"`
	Rain        []float64 `json:"rain"`
}

func main() {
	var weather weather
	var baseURL string = "https://api.open-meteo.com/v1/forecast?"
	var queryURL string = "&hourly=temperature_2m,rain,windspeed_10m"

	var coord coordinates
	coord.latitude, coord.longitude = 52.52, 13.41

	var URL string = buildURL(baseURL, queryURL, coord)

	response, err := http.Get(URL)

	if err != nil {
		fmt.Println("Im so bad at programming")
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body) // response body is []byte

	var result result
	json.Unmarshal(body, &result)

	weather = &result

	var info weatherInfo = CoallesceWeatherInfo(weather)
	fmt.Println("Current Temperature: " + fmt.Sprintf("%v", info.temperature))
	fmt.Println("Current Wind Speed: " + fmt.Sprintf("%v", info.wind))
	fmt.Println("Rain chance(%): " + fmt.Sprintf("%v", info.rain))

}
