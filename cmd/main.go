package main

import (
	"log"
	"net/http"

	"github.com/renanmav/GoExpert-CEPTemperature-GCR/config"
	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/delivery"
	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/third_party_api"
	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	cepAPI := third_party_api.NewCepApi()
	weatherAPI := third_party_api.NewWeatherApi(cfg.WeatherAPIKey)

	weatherByCepUseCase := usecase.NewWeatherByCepUseCase(cepAPI, weatherAPI)

	handler := delivery.NewHTTPHandler(weatherByCepUseCase)
	http.HandleFunc("/weather", handler.GetWeather)

	log.Printf("Server listening on port %s", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
