package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type article struct {
	ID        int
	Name      string
	CreatedAt time.Time
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

func main() {
	http.HandleFunc("/articles", getArticle)
	fmt.Println("Starting Web Server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
