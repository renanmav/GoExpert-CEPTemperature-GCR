package usecase

import (
	"fmt"

	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/third_party_api"
)

type WeatherByCepInput struct {
	CEP string
}

type WeatherByCepOutput struct {
	City       string `json:"city"`
	Celsius    string `json:"temp_C"`
	Fahrenheit string `json:"temp_F"`
	Kelvin     string `json:"temp_K"`
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

func (s *WeatherByCepUseCase) GetWeatherByCEP(input WeatherByCepInput) (*WeatherByCepOutput, error) {
	location, err := s.cepGetter.GetLocationByCEP(input.CEP)
	if err != nil {
		return nil, err
	}

	C, err := s.weatherGetter.GetWeatherByCity(location.City)
	if err != nil {
		return nil, err
	}

	F := C*1.8 + 32
	K := C + 273

	return &WeatherByCepOutput{
		City:       location.City,
		Celsius:    fmt.Sprintf("%.2f", C),
		Fahrenheit: fmt.Sprintf("%.2f", F),
		Kelvin:     fmt.Sprintf("%.2f", K),
	}, nil
}
