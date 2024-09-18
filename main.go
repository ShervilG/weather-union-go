package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	WEATHER_BASE_URL = "https://www.weatherunion.com/gw/weather/external/v0"
	API_KEY          = os.Getenv("WEATHER_UNION_API_KEY")
	LOCALITY_ID      = "ZWL007653"
)

type WeatherData struct {
	Temperature      float64 `json:"temperature"`
	Humidity         float64 `json:"humidity"`
	WindSpeed        float64 `json:"wind_speed"`
	WindDirection    float64 `json:"wind_direction"`
	RainIntensity    float64 `json:"rain_intensity"`
	RainAccumulation float64 `json:"rain_accumulation"`
}

type WeatherResponse struct {
	Status              string       `json:"status"`  // non-nullable
	Message             string       `json:"message"` // non-nullable
	DeviceType          *int         `json:"device_type,omitempty"`
	LocalityWeatherData *WeatherData `json:"locality_weather_data,omitempty"`
}

var httpClient *http.Client

func main() {
	httpClient = &http.Client{}

	url := fmt.Sprintf("%v/get_locality_weather_data?locality_id=%v", WEATHER_BASE_URL, LOCALITY_ID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-zomato-api-key", API_KEY)

	res, _ := httpClient.Do(req)
	bodyBytes, _ := io.ReadAll(res.Body)

	fmt.Println(string(bodyBytes))
}
