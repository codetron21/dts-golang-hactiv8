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
	"time"
)

const (
	fileName = "status.json"
)

var data File

type File struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	go loopCreateReadData()
	http.HandleFunc("/", handleRoute)
	http.ListenAndServe(":8080", nil)
}

func loopCreateReadData() {
	for {
		fmt.Println("run")

		// generate data
		data = File{
			Status: Status{
				Water: getRandomNumber(),
				Wind:  getRandomNumber(),
			},
		}

		// write json file
		err := createJsonFile(fileName, data)
		if err != nil {
			println("error create json file", err)
			return
		}

		// read json file
		err = readJsonFile(fileName, &data)
		if err != nil {
			println("error read json file", err)
			return
		}

		time.Sleep(3 * time.Second)
	}
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	// check http method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusInternalServerError)
		return
	}

	// read html file
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// templating
	waterStatus, windStatus := data.Status.checkStatus()

	err = tmpl.Execute(w, map[string]interface{}{
		"water":       data.Status.Water,
		"wind":        data.Status.Wind,
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

func createJsonFile(fileName string, data File) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("error create file", err)
		return err
	}

	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		fmt.Println("error write file", err)
		return err
	}

	return nil
}

func readJsonFile(fileName string, data *File) error {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("error read file", err)
		return err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error read file json", err)
		return err
	}

	jsonFile.Close()

	err = json.Unmarshal(byteValue, data)
	if err != nil {
		fmt.Println("error read file unmarshal json")
		return err
	}

	return nil
}

func getRandomNumber() int {
	return rand.Intn(100) + 1
}
