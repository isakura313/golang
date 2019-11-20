package main

import (
	"database/sql"
	"encoding/json" // Для корректной работы с json
	"fmt"           // библиотека для вывода
	"log"           // для логирования
	"net/http"      // и для обработки http запросов

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux" //http роутер  и диспатчер
)

type Article struct {
	Title   string `json:"Title"`
	ID      int    `json: id`
	Desc    string `json:"desc"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	articles := Articles{
		Article{Title: "О дельфинах", ID: 1, Desc: "дельфинчики клевые", Content: "Дельфи́ны — водные млекопитающие отряда китообразных, принадлежащие либо к семейству дельфиновых — морские, либо к надсемейству речных дельфинов — пресноводные."},
		Article{Title: "О пингвинах", ID: 2, Desc: "Пингвины классные", Content: "Пингви́новые, или пингви́ны, — семейство нелетающих морских птиц, единственное современное в отряде пингвинообра́зных."},
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

type User struct {
	id   int
	Name string
}

func main() {
	// handleRequests()
	// sqlConnString := getConnString()
	db, err := sql.Open("mysql", "pavel:@tcp(127.0.0.1:3306)/testdb")
	// db, err := sql.Open("mysql", "pavel:@/root1")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	stmtIns, err := db.Prepare("INSERT INTO shop VALUES ('Pavel',1)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	rows, err := db.Query("SELECT Name FROM shop")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		var user User

		err = rows.Scan(&user.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		fmt.Println(user.Name)
	}
	defer rows.Close()
}
