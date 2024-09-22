package third_party_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/entity"
)

type CepApi struct{}

func NewCepApi() CepApiInterface {
	return &CepApi{}
}

func (api *CepApi) GetLocationByCEP(cep string) (*entity.Location, error) {
	fmt.Println("Getting location by CEP: ", cep)

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		CEP        string `json:"cep"`
		Localidade string `json:"localidade"`
		UF         string `json:"uf"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Localidade == "" {
		return nil, errors.New("location not found for CEP: " + cep)
	}

	fmt.Println("Location found: ", result)

	return &entity.Location{
		CEP:   result.CEP,
		City:  result.Localidade,
		State: result.UF,
	}, nil
}
