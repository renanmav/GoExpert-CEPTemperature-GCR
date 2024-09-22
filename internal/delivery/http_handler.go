package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/usecase"
)

type HTTPHandler struct {
	weatherByCepUseCase *usecase.WeatherByCepUseCase
}

func NewHTTPHandler(weatherByCepUseCase *usecase.WeatherByCepUseCase) *HTTPHandler {
	return &HTTPHandler{weatherByCepUseCase: weatherByCepUseCase}
}

func (h *HTTPHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP parameter is required", http.StatusBadRequest)
		return
	}

	weather, err := h.weatherByCepUseCase.GetWeatherByCEP(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
