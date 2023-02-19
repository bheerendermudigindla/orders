// Go code
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Details of order struct to save orders
type Order struct {
	ID           string  `json:"id"`
	Status       string  `json:"status"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total"`
	CurrencyUnit string  `json:"currencyUnit"`
}

// Details Item struct to save items in order
type Item struct {
	ID          string  `json:"id"`
	Order_id    int     `json:"order_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/orders", getOrders)
	http.HandleFunc("/order", createOrder)
	http.HandleFunc("/order/update", updateOrder)
	log.Fatal(http.ListenAndServe(":8080", nil))
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
	dbname := "mysql_db"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname))
	log.Println("Error: ", err)

	handleErr(err)
	return db
}

// Creating new order
func createOrder(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var order Order
		json.Unmarshal(body, &order)

		// Insert order into db
		stmt, err := db.Prepare("INSERT INTO order_data (id, status, total, currencyUnit) VALUES (?, ?, ?, ?)")
		handleErr(err)
		_, err = stmt.Exec(order.ID, order.Status, order.Total, order.CurrencyUnit)
		handleErr(err)

		// Insert items into db
		for _, item := range order.Items {
			stmt, err := db.Prepare("INSERT INTO order_items (id, order_id, description, price, quantity) VALUES (?, ?, ?, ?, ?)")
			handleErr(err)
			_, err = stmt.Exec(item.ID, item.Order_id, item.Description, item.Price, item.Quantity)
			handleErr(err)
		}
		fmt.Fprintf(w, "New Order Created Successfully")
	}
	defer db.Close()
}

// Get Orders details from database
func getOrders(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	log.Println("method:: ", r.Method)
	if r.Method == "GET" {
		log.Println("getorders")
		var query strings.Builder
		query.WriteString("SELECT * FROM order_data")

		// Get query params details
		params := r.URL.Query()
		// Filter orders
		if len(params) > 0 {
			query.WriteString(" WHERE ")
			for key, value := range params {
				query.WriteString(key)
				query.WriteString(" = ")
				query.WriteString("'")
				query.WriteString(value[0])
				query.WriteString("'")
				query.WriteString(" AND ")
			}
			// Remove last AND
			queryStr := query.String()
			queryStr = queryStr[:len(queryStr)-4]
			// Sort orders
			queryStr = queryStr + " ORDER BY id"

			log.Println("Getorderlist query:: ", queryStr)

			// Execute query
			rows, err := db.Query(queryStr)
			handleErr(err)
			// Fetch orders
			var orders []Order
			for rows.Next() {
				var order Order
				err = rows.Scan(&order.ID, &order.Status, &order.Total, &order.CurrencyUnit)
				handleErr(err)
				// Fetch items
				itemsQuery := "SELECT * FROM order_items WHERE order_id=" + order.ID
				log.Println("Getorderlist itemsQuery:: ", itemsQuery)
				itemRows, err := db.Query(itemsQuery)
				handleErr(err)
				var items []Item
				for itemRows.Next() {
					var item Item
					err = itemRows.Scan(&item.ID, &item.Order_id, &item.Description, &item.Price, &item.Quantity)
					handleErr(err)
					// Append item to order's items
					items = append(items, item)
				}
				order.Items = items
				// Append order to orders
				orders = append(orders, order)
			}
			// Set response header
			w.Header().Set("Content-Type", "application/json")
			// Convert orders to JSON
			json.NewEncoder(w).Encode(orders)
		}
	}
	defer db.Close()
}

// Update existing Order in database
func updateOrder(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var order Order
		json.Unmarshal(body, &order)
		// Update order in database
		stmt, err := db.Prepare("UPDATE order_data SET status=? WHERE id=?")
		handleErr(err)
		_, err = stmt.Exec(order.Status, order.ID)
		handleErr(err)
		fmt.Fprintf(w, "Order Updated Successfully")
	}
	defer db.Close()
}
