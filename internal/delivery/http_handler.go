package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/usecase"
)

var (
	ErrInvalidCEP = errors.New("invalid zipcode")
	ErrNotFound   = errors.New("can not find zipcode")
)

type HTTPHandler struct {
	weatherByCepUseCase *usecase.WeatherByCepUseCase
}

func NewHTTPHandler(weatherByCepUseCase *usecase.WeatherByCepUseCase) *HTTPHandler {
	return &HTTPHandler{weatherByCepUseCase: weatherByCepUseCase}
}

func (h *HTTPHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	if err := h.ValidateCEP(cep); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	weather, err := h.weatherByCepUseCase.GetWeatherByCEP(usecase.WeatherByCepInput{CEP: cep})
	if err != nil {
		wrappedErr := fmt.Errorf("%w: %v", ErrNotFound, err)
		http.Error(w, wrappedErr.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func (h *HTTPHandler) ValidateCEP(cep string) error {
	cepRegex := `^(\d{5}-?\d{3})$`
	match, err := regexp.MatchString(cepRegex, cep)
	if err != nil || !match {
		return ErrInvalidCEP
	}
	return nil
}
