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
	//классов нет но есть struct и, и у них есть имя и поля, в которые можем вписать поля нашего объекта, используя struct как Конструктор
}

type Articles []Article // объявляем массив наших структов, что бы можно было складывать наши статьи

func allArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// для того, что наше приложение могло отдать GET - запрос, нам нужно
	// установить заголовки. Подробнее про них можно прочитать
	//https://developer.mozilla.org/ru/docs/Web/HTTP/%D0%97%D0%B0%D0%B3%D0%BE%D0%BB%D0%BE%D0%B2%D0%BA%D0%B8
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//если вкратче, данный заголовок открывает всем доступ к данному запросу.
	// используется только в учебных целях, в профессиональной продакшене это, конечно, возможно стоит  поменять
	sql := "SELECT * from articles" //выбрать все из таблицы articles
	rows, err := getJSON(sql)
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
	log.Println(http.ListenAndServe(":8810", myRouter))
}

func getJSON(sqlString string) (string, error) {
	db, err := sql.Open("mysql", "pavel:@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Println(err.Error()) 
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
