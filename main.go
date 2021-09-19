package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type article struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

type PeopleSwapi struct {
	BirthYear string    `json:"birth_year"`
	EyeColor  string    `json:"eye_color"`
	Films     []string  `json:"films"`
	Gender    string    `json:"gender"`
	HairColor string    `json:"hair_color"`
	Height    string    `json:"height"`
	Homeworld string    `json:"homeworld"`
	Mass      string    `json:"mass"`
	Name      string    `json:"name"`
	SkinColor string    `json:"skin_color"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	Species   []string  `json:"species"`
	Starships []string  `json:"starships"`
	URL       string    `json:"url"`
	Vehicles  []string  `json:"vehicles"`
}

var articles = []article{
	{
		ID:        1,
		Name:      "Article_1",
		CreatedAt: time.Now(),
	},
	{
		ID:        2,
		Name:      "Article_2",
		CreatedAt: time.Now().AddDate(0, 0, -2),
	},
}

func getArticle(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if req.Method == "GET" {
		result, err := json.Marshal(articles)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Write(result)
		return
	}
	http.Error(res, "Method Not Allowed", http.StatusBadRequest)
	return
}

func getSWPeople(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if req.Method == "GET" {
		result, err := http.Get("https://swapi.dev/api/people/1/")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		responseData, _ := ioutil.ReadAll(result.Body)
		defer result.Body.Close()

		var peopleSW PeopleSwapi
		json.Unmarshal(responseData, &peopleSW)
		//if u wanna print at CLI
		//fmt.Println(responseData)

		resultJson, err := json.Marshal(peopleSW)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Write(resultJson)
		return

	}
	http.Error(res, "Method Not Allowed", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/articles", getArticle)
	http.HandleFunc("/people", getSWPeople)
	fmt.Println("Starting Web Server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
