package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
)

type File struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusInternalServerError)
			return
		}

		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonFile, err := loadJsonFile("status.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonFile.Close()

		var data File

		json.Unmarshal(byteValue, &data)

		data.Status.Water = getRandomNumber()
		data.Status.Wind = getRandomNumber()
		waterStatus, windStatus := data.Status.checkStatus()

		fmt.Println(data)

		err = tmpl.Execute(w, map[string]interface{}{
			"water":       data.Status.Water,
			"wind":        data.Status.Wind,
			"waterStatus": waterStatus,
			"windStatus":  windStatus,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe("localhost:8080", nil)
}

func (s Status) checkStatus() (waterStatus string, windStatus string) {
	if s.Water <= 5 {
		waterStatus = "aman"
	} else if s.Water > 5 && s.Water < 9 {
		waterStatus = "siaga"
	} else {
		waterStatus = "bahaya"
	}

	if s.Wind <= 6 {
		windStatus = "aman"
	} else if s.Wind > 6 && s.Wind < 16 {
		windStatus = "siaga"
	} else {
		windStatus = "bahaya"
	}

	return
}

func loadJsonFile(fileName string) (*os.File, error) {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	return jsonFile, nil
}

func getRandomNumber() int {
	return rand.Intn(100) + 1
}
