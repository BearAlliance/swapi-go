package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func swapiRequest(thing string, number int) []byte {
	baseUrl := "https://swapi.dev/api"
	url := fmt.Sprintf("%s/%s/%d", baseUrl, thing, number)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Fetch Error!")
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println("io Read error!")
	}

	return body
}

type Film struct {
	Title string
	Eipsode_id int
	Opening_crawl string
	Director string
	Producer string
	Release_date string
	Characters []string
	Planets []string
	Starships []string
	Vehicles []string
	Species []string
	Created string
	Edited string
	Url string
}

func GetFilm(number int) Film {
	body := swapiRequest("films", number)

	var f Film
	jsonErr := json.Unmarshal(body, &f)

	if jsonErr != nil {
		fmt.Print("\nJSON unmarshall error: ")
		fmt.Print(jsonErr)
	}

	return f
}