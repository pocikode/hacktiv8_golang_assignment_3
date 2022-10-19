package main

import (
	"assignment-3/services"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	go services.UpdateWeather()

	http.HandleFunc("/", home)
	http.HandleFunc("/api/weather", getWeather)
	http.ListenAndServe(":9000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl, errTmpl := template.ParseFiles("static/index.html")
	if errTmpl != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": errTmpl.Error(),
		})
		return
	}

	tpl.Execute(w, r)
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	weatherFile, err := os.ReadFile("data.json")
	if err != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	var weatherData services.WeatherData
	err = json.Unmarshal(weatherFile, &weatherData)
	if err != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	water := weatherData.Status.Water
	wind := weatherData.Status.Wind
	var waterStatus, windStatus string

	if water <= 5 {
		waterStatus = "Aman"
	} else if water >= 6 && water <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}

	if wind <= 6 {
		windStatus = "Aman"
	} else if wind >= 7 && wind <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}

	weatherResult := services.WeatherJsonResponse{}

	weatherResult.Wind.Value = wind
	weatherResult.Wind.ValueText = fmt.Sprintf("%d m/s", wind)
	weatherResult.Wind.Status = windStatus

	weatherResult.Water.Value = water
	weatherResult.Water.ValueText = fmt.Sprintf("%d m", water)
	weatherResult.Water.Status = waterStatus

	writeJsonResponse(w, http.StatusOK, weatherResult)
	return
}

func writeJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
