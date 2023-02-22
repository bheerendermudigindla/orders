package main

import (
	"Golang_crud/routes"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/orders", ordersHandler)
	http.ListenAndServe(":8080", nil)
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		routes.GetOrders(w, r)
	case "POST":
		routes.CreateOrder(w, r)
	case "PUT":
		routes.UpdateOrder(w, r)
	default:
		msg := "Method " + r.Method + " Not Allowed!"
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}
}
