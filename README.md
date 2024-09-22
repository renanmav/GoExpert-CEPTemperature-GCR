# GoExpert-CEPTemperature-GCR

This project is a Go-based web service that provides temperature information for a given CEP (Postal Code) in Brazil. It uses the ViaCEP API to get location information and the WeatherAPI to fetch temperature data. The service is designed to be deployed on Google Cloud Run.

## Features

- Fetch location information using CEP (Brazilian Postal Code)
- Retrieve current temperature for the location
- Convert temperature to Celsius, Fahrenheit, and Kelvin
- Configurable using environment variables
- Designed with Clean Architecture principles

## Prerequisites

- Go 1.19 or later
- WeatherAPI API key (sign up at https://www.weatherapi.com)

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/renanmav/GoExpert-CEPTemperature-GCR.git
   cd GoExpert-CEPTemperature-GCR
   ```

1. Create a `.env` file from the provided `.env.example`:
   ```
   cp .env.example .env
   ```

1. Edit the `.env` file and set your WeatherAPI key:
   ```
   WEATHER_API_KEY=your_api_key_here
   ```

1. Install dependencies:
   ```
   go mod tidy
   ```

## Running the Application

To run the application locally:

```
go run cmd/main.go
```

The server will start on the port specified in your `.env` file (default is 8080).

## API Usage

To get weather information for a CEP, make a GET request to the `/weather` endpoint with the `cep` query parameter:

```bash
curl http://localhost:8080/weather?cep=01001000
```

The response will be a JSON object containing the city name and temperature in Celsius, Fahrenheit, and Kelvin.

## Deployment to Google Cloud Run

1. Build the Docker image:
   ```
   docker build -t gcr.io/your-project-id/cep-temperature-service .
   ```

2. Push the image to Google Container Registry:
   ```
   docker push gcr.io/your-project-id/cep-temperature-service
   ```

3. Deploy to Cloud Run:
   ```
   gcloud run deploy --image gcr.io/your-project-id/cep-temperature-service --platform managed
   ```

Make sure to set the `WEATHER_API_KEY` environment variable in your Cloud Run configuration.
