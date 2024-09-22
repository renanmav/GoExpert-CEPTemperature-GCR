package third_party_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherApi struct {
	apiKey string
}

func NewWeatherApi(apiKey string) WeatherApiInterface {
	return &WeatherApi{apiKey: apiKey}
}

func (api *WeatherApi) GetWeatherByCoordinates(lat, lon float64) (float64, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%.6f,%.6f", api.apiKey, lat, lon)
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

	return result.Current.TempC, nil
}
