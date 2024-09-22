package third_party_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WeatherApi struct {
	apiKey string
}

func NewWeatherApi(apiKey string) WeatherApiInterface {
	return &WeatherApi{apiKey: apiKey}
}

func (api *WeatherApi) GetWeatherByCity(city string) (float64, error) {
	fmt.Println("Getting weather by city: ", city)

	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", api.apiKey, encodedCity)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	fmt.Println("Weather found: ", result)

	return result.Current.TempC, nil
}
