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

	//fmt.Println(URL)
	if err != nil {
		fmt.Println("Im so bad at programming")
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body) // response body is []byte

	var result result
	//fmt.Println(string(body))
	json.Unmarshal(body, &result)

	weather = &result
	fmt.Println("Current Temperature: " + fmt.Sprintf("%v", weather.CurrentTemperature()))
	fmt.Println("Current Wind Speed: " + fmt.Sprintf("%v", weather.CurrentWind()))

	if weather.CurrentRain() == 0.00 {
		fmt.Println("It is not raining")
	} else {
		fmt.Println("Rain chance(%): " + fmt.Sprintf("%v", weather.CurrentRain()))
	}

}
