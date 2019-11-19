package main

import (
	"encoding/json" // Для корректной работы с json
	"fmt"           // библиотека для вывода
	"log"           // для логирования
	"net/http"      // и для обработки http запросов

	"github.com/gorilla/mux" //http роутер  и диспатчер
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "О дельфинах", Desc: "дельфинчики клевые", Content: "Дельфинчики просто супер"},
	}

	fmt.Println("Endpoint Hit:All articles")
	json.NewEncoder(w).Encode(articles)
}
func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>test Post endpoint POST </h1>")

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Наш сервер начал работу </h1>")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("Post")
	log.Fatal(http.ListenAndServe(":8801", myRouter))
}

func main() {
	handleRequests()
}
