package routes

import (
	"database/sql"
	"fmt"
	"log"
)

type Order struct {
	ID           string      `json:"id"`
	Status       string      `json:"status"`
	Items        []OrderItem `json:"items"`
	Total        float64     `json:"total"`
	CurrencyUnit string      `json:"currencyUnit"`
}

type OrderItem struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

// Handle Error
func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to mysql successfully")
	}
}

// Connection to mysql DB
func dbConn() (db *sql.DB) {
	host := "localhost"
	port := 3306
	user := "root"
	password := "User@123"
	dbname := "task"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname))
	log.Println("Error: ", err)

	handleErr(err)
	return db
}
