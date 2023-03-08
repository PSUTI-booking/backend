package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Sbook struct {
	Id          int    `json:"id"`
	Class       string `json:"class"`
	Description string `json:"description"`
	Users       string `json:"users"`
}

func main() {
	//открываем базу данных
	connStr := "user=postgres password=Amogus228! dbname=bookingpsuti sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	//делаем запрос через db.query
	rows, err := db.Query("select * from bookingpsuti")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	book := []Sbook{} //создаем массив/слайс структур , в него будем всё записывать

	for rows.Next() { //с помощью метода rows.Next() мы можем последовательно перебрать все строки в полученном наборе
		boo := Sbook{}
		//Тип rows определяет метод Scan, с помощью которого можно считать все полученные данные в переменные
		err := rows.Scan(&boo.Id, &boo.Class, &boo.Description, &boo.Users)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//считываем данные в структуру Employer и затем добавляем ее в срез
		book = append(book, boo) // добавляем в слайс переменные , созданные из данных таблицы
	}

	e := echo.New() //инстанция

	e.GET("/booking", func(c echo.Context) error {
		return c.JSONPretty(200, book, " ")
	})

	defer e.Logger.Fatal(e.Start(":1323"))
}
