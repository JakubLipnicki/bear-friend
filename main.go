package main

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	NWSAPIBase = "https://api.weather.gov"
	UserAgent  = "weather-app/1.0"
)

type PointsResponse struct {
	Properties struct {
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

type ForecastResponse struct {
	Properties struct {
		Periods []ForecastPeriod `json:"periods"`
	} `json:"properties"`
}

type ForecastPeriod struct {
	Name             string `json:"name"`
	Temperature      int    `json:"temperature"`
	TemperatureUnit  string `json:"temperatureUnit"`
	WindSpeed        string `json:"windSpeed"`
	WindDirection    string `json:"windDirection"`
	DetailedForecast string `json:"detailedForecast"`
}

type AlertResponse struct {
	Features []AlertFeature `json:"features"`
}

type AlertFeature struct {
	Properies AlertProperties `json:"properties"`
}

type AlertProperties struct {
	Event       string `json:"event"`
	AreaDesc    string `json:areaDesc`
	Severity    string `json:"severity`
	Descritpion string `json:"description"`
	Instruction string `json:"Instruction"`
}

func main() {
	fmt.Println("test")
}
