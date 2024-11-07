package usecase_test

import (
	"errors"
	"testing"

	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/entity"
	"github.com/renanmav/GoExpert-CEPTemperature-GCR/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock CepAPI
type MockCepAPI struct {
	mock.Mock
}

func (m *MockCepAPI) GetLocationByCEP(cep string) (*entity.Location, error) {
	args := m.Called(cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Location), args.Error(1)
}

// Mock WeatherAPI
type MockWeatherAPI struct {
	mock.Mock
}

func (m *MockWeatherAPI) GetWeatherByCity(city string) (float64, error) {
	args := m.Called(city)
	return args.Get(0).(float64), args.Error(1)
}

func TestGetWeatherByCEP(t *testing.T) {
	// Common errors
	errCepNotFound := errors.New("CEP not found")
	errCepInvalid := errors.New("invalid CEP format")
	errWeatherAPIDown := errors.New("weather API is unavailable")
	errCityNotFound := errors.New("city not found in weather database")

	type TestCase struct {
		name          string
		input         usecase.WeatherByCepInput
		mockLocation  *entity.Location
		mockTemp      float64
		setupMocks    func(*MockCepAPI, *MockWeatherAPI, usecase.WeatherByCepInput, *entity.Location, float64)
		expectedError error
		expected      *usecase.WeatherByCepOutput
	}

	tests := []TestCase{
		{
			name: "Success case - room temperature",
			input: usecase.WeatherByCepInput{
				CEP: "12345-678",
			},
			mockLocation: &entity.Location{
				City: "São Paulo",
			},
			mockTemp: 25.0,
			setupMocks: func(mc *MockCepAPI, mw *MockWeatherAPI, input usecase.WeatherByCepInput, loc *entity.Location, temp float64) {
				mc.On("GetLocationByCEP", input.CEP).Return(loc, nil)
				mw.On("GetWeatherByCity", loc.City).Return(temp, nil)
			},
			expected: &usecase.WeatherByCepOutput{
				City:       "São Paulo",
				Celsius:    "25.00",
				Fahrenheit: "77.00",
				Kelvin:     "298.00",
			},
		},
		{
			name: "Success case - very hot temperature",
			input: usecase.WeatherByCepInput{
				CEP: "12345-679",
			},
			mockLocation: &entity.Location{
				City: "Death Valley",
			},
			mockTemp: 50.0,
			setupMocks: func(mc *MockCepAPI, mw *MockWeatherAPI, input usecase.WeatherByCepInput, loc *entity.Location, temp float64) {
				mc.On("GetLocationByCEP", input.CEP).Return(loc, nil)
				mw.On("GetWeatherByCity", loc.City).Return(temp, nil)
			},
			expected: &usecase.WeatherByCepOutput{
				City:       "Death Valley",
				Celsius:    "50.00",
				Fahrenheit: "122.00",
				Kelvin:     "323.00",
			},
		},
		{
			name: "Success case - very cold temperature",
			input: usecase.WeatherByCepInput{
				CEP: "12345-680",
			},
			mockLocation: &entity.Location{
				City: "Antarctica",
			},
			mockTemp: -50.0,
			setupMocks: func(mc *MockCepAPI, mw *MockWeatherAPI, input usecase.WeatherByCepInput, loc *entity.Location, temp float64) {
				mc.On("GetLocationByCEP", input.CEP).Return(loc, nil)
				mw.On("GetWeatherByCity", loc.City).Return(temp, nil)
			},
			expected: &usecase.WeatherByCepOutput{
				City:       "Antarctica",
				Celsius:    "-50.00",
				Fahrenheit: "-58.00",
				Kelvin:     "223.00",
			},
		},
		{
			name: "Error - CEP not found",
			input: usecase.WeatherByCepInput{
				CEP: "00000-000",
			},
			setupMocks: func(mc *MockCepAPI, mw *MockWeatherAPI, input usecase.WeatherByCepInput, loc *entity.Location, temp float64) {
				mc.On("GetLocationByCEP", input.CEP).Return(nil, errCepNotFound)
			},
			expectedError: errCepNotFound,
		},
		{
			name: "Error - Invalid CEP format",
			input: usecase.WeatherByCepInput{
				CEP: "invalid",
			},
			setupMocks: func(mc *MockCepAPI, mw *MockWeatherAPI, input usecase.WeatherByCepInput, loc *entity.Location, temp float64) {
				mc.On("GetLocationByCEP", input.CEP).Return(nil, errCepInvalid)
			},
			expectedError: errCepInvalid,
		},
		{
			name: "Error - Weather API down",
			input: usecase.WeatherByCepInput{
				CEP: "12345-681",
			},
			mockLocation: &entity.Location{
				City: "Rio de Janeiro",
			},
			setupMocks: func(mc *MockCepAPI, mw *MockWeatherAPI, input usecase.WeatherByCepInput, loc *entity.Location, temp float64) {
				mc.On("GetLocationByCEP", input.CEP).Return(loc, nil)
				mw.On("GetWeatherByCity", loc.City).Return(0.0, errWeatherAPIDown)
			},
			expectedError: errWeatherAPIDown,
		},
		{
			name: "Error - City not found in weather database",
			input: usecase.WeatherByCepInput{
				CEP: "12345-682",
			},
			mockLocation: &entity.Location{
				City: "Unknown City",
			},
			setupMocks: func(mc *MockCepAPI, mw *MockWeatherAPI, input usecase.WeatherByCepInput, loc *entity.Location, temp float64) {
				mc.On("GetLocationByCEP", input.CEP).Return(loc, nil)
				mw.On("GetWeatherByCity", loc.City).Return(0.0, errCityNotFound)
			},
			expectedError: errCityNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockCepAPI := new(MockCepAPI)
			mockWeatherAPI := new(MockWeatherAPI)

			// Set expectations using the setup function
			tt.setupMocks(mockCepAPI, mockWeatherAPI, tt.input, tt.mockLocation, tt.mockTemp)

			// Create use case with mocks
			useCase := usecase.NewWeatherByCepUseCase(mockCepAPI, mockWeatherAPI)

			// Execute
			result, err := useCase.GetWeatherByCEP(tt.input)

			// Assert
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}

			// Verify mock expectations
			mockCepAPI.AssertExpectations(t)
			mockWeatherAPI.AssertExpectations(t)
		})
	}
}

func TestNewWeatherByCepUseCase(t *testing.T) {
	mockCepAPI := new(MockCepAPI)
	mockWeatherAPI := new(MockWeatherAPI)

	useCase := usecase.NewWeatherByCepUseCase(mockCepAPI, mockWeatherAPI)

	assert.NotNil(t, useCase)
}
