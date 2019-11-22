package main

import (
	"database/sql"  // основной плагин для использования sql
	"encoding/json" // Для корректной работы с json
	"fmt"           // библиотека для вывода
	"log"           // для логирования
	"net/http"      // и для обработки http запросов

	_ "github.com/go-sql-driver/mysql" //драйвер для работы нашего sql
	"github.com/gorilla/mux"           //http роутер  и диспатчер
)

type Article struct {
	ID      int    `json: id`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"Content"`
	//классов нет но есть struct и они классные
}

type Articles []Article // объявляем массив наших структов

func allArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sql := "SELECT * from shop" //сложнейший sql не поддающийся никакому анализу
	rows, err := getJSON(sql)
	log.Println(rows) // нужно для отладки и работы. Но вообще можно и удалить если что
	log.Println(err)
	articles := rows
	fmt.Println("Endpoint Hit:All articles")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Элвис покинул здание</h1>")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8801", myRouter))
}

func getJSON(sqlString string) (string, error) {
	db, err := sql.Open("mysql", "pavel:@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error()) // panic относительно бесполезен. В реальности вам нужен другой перехватчик ошибок
	}
	defer db.Close()

	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func main() {
	handleRequests()
}
