package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		var order Order
		rows.Scan(&order.ID, &order.Status, &order.Total, &order.CurrencyUnit)
		log.Println("order:: ", order)

		Items := []OrderItem{}
		query := "SELECT id, description, price, quantity FROM order_items where order_id = '" + order.ID + "'"
		log.Println("query: ", query)
		data, err1 := db.Query(query)
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer data.Close()
		for data.Next() {
			var item OrderItem
			err2 := data.Scan(&item.ID, &item.Description, &item.Price, &item.Quantity)
			if err2 != nil {
				log.Println("Error at item scan:", err2)
			}
			Items = append(Items, item)
		}
		order.Items = Items
		log.Println("orderdata:: ", order)
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
