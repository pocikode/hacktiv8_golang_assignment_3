package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type WeatherData struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
}

type WeatherJsonResponse struct {
	Water WeatherResult `json:"water"`
	Wind  WeatherResult `json:"wind"`
}

type WeatherResult struct {
	Value     int    `json:"value"`
	ValueText string `json:"value_text"`
	Status    string `json:"status"`
}

func UpdateWeather() {
	max := 100
	min := 1
	for {
		fmt.Println("Running Update Weather per 15 Second")
		rand.Seed(time.Now().UnixNano())
		statusWeather := WeatherData{}
		statusWeather.Status.Water = rand.Intn(max-min) + min
		statusWeather.Status.Wind = rand.Intn(max-min) + min

		dataWeather, err := json.MarshalIndent(statusWeather, "", "\t")
		if err != nil {
			fmt.Println(err)
		}

		errs := os.WriteFile("data.json", dataWeather, 0644)
		if errs != nil {
			fmt.Println(errs)
		}
		time.Sleep(15 * time.Second)
	}
}
