package third_party_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/entity"
)

type CepApi struct{}

func NewCepApi() CepApiInterface {
	return &CepApi{}
}

func (api *CepApi) GetLocationByCEP(cep string) (*entity.Location, error) {
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

	// Note: ViaCEP doesn't provide latitude and longitude.
	// You might need to use another API or service to get these coordinates.
	return &entity.Location{
		CEP:   result.CEP,
		City:  result.Localidade,
		State: result.UF,
	}, nil
}
