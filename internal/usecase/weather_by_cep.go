package usecase

import (
	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/entity"
	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/third_party_api"
)

type WeatherByCepInput struct {
	CEP string
}

type WeatherByCepOutput struct {
	City       string  `json:"city"`
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Kelvin     float64 `json:"kelvin"`
}

type WeatherByCepUseCase struct {
	cepGetter     third_party_api.CepApiInterface
	weatherGetter third_party_api.WeatherApiInterface
}

func NewWeatherByCepUseCase(cepGetter third_party_api.CepApiInterface, weatherGetter third_party_api.WeatherApiInterface) *WeatherByCepUseCase {
	return &WeatherByCepUseCase{
		cepGetter:     cepGetter,
		weatherGetter: weatherGetter,
	}
}

func (s *WeatherByCepUseCase) GetWeatherByCEP(cep string) (*entity.Weather, error) {
	location, err := s.cepGetter.GetLocationByCEP(cep)
	if err != nil {
		return nil, err
	}

	tempCelsius, err := s.weatherGetter.GetWeatherByCoordinates(location.Latitude, location.Longitude)
	if err != nil {
		return nil, err
	}

	tempFahrenheit := tempCelsius*1.8 + 32
	tempKelvin := tempCelsius + 273.15

	return &entity.Weather{
		City:       location.City,
		Celsius:    tempCelsius,
		Fahrenheit: tempFahrenheit,
		Kelvin:     tempKelvin,
	}, nil
}
