package third_party_api

import (
	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/entity"
)

type CepApiInterface interface {
	GetLocationByCEP(cep string) (*entity.Location, error)
}

type WeatherApiInterface interface {
	GetWeatherByCoordinates(lat, lon float64) (tempCelsius float64, err error)
}
