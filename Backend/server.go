// Backend Server

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Order struct {
	OrderName       string  `json:"orderName"`
	ProductName     string  `json:"productName"`
	CustomerName    string  `json:"customerName"`
	CustomerCompany string  `json:"customerCompany"`
	OrderDate       string  `json:"orderDate"`
	DeliveredAmount float32 `json:"deliveredAmount"`
	TotalAmount     float32 `json:"totalAmount"`
}

type OrderArray struct {
	orders []Order `json:"orders"`
}

func main() {
	http.HandleFunc("/", filter)
	http.ListenAndServe(":8090", nil)
}

func filter(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// process data pass from frent end
	// improve below code by set json request from front end
	data := strings.Split(string(body), ",")
	search := data[0][2:(len(data[0]) - 1)]
	startDate := data[1][1:(len(data[1]) - 1)]
	endDate := data[2][1:(len(data[2]) - 2)]

	dataRespond := retriveData(search, startDate, endDate)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dataRespond); err != nil {
		panic(err)
	}

}

func retriveData(search string, startDate string, endDate string) []Order {

	// read db connection credential from config file
	postgresSQL, err := os.ReadFile("./config/postgresql.config")
	checkErr(err)

	mongoDB, err := os.ReadFile("./config/mongodb.config")
	checkErr(err)

	// MongoDB connection
	uri := string(mongoDB)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	// retrive mongodb data
	customersCollection := client.Database("OrderBrowser").Collection("Customers")
	cus_CompaniesCollection := client.Database("OrderBrowser").Collection("Customer_companies")

	// PostgreDB connection
	postgreDB, err := sql.Open("postgres", string(postgresSQL))
	checkErr(err)

	// set query string for sql database
	// code quality improve: put these query string into an independent file and read string from that file
	queryString1 := "select C.customer_id, C.order_name, C.product, C.created_at, D.delivered_quantity*C.PRICE_PER_UNIT AS Deliveried_Amount, C.price_per_unit*C.quantity AS Total_Amount "
	queryString2 := "from (select A.customer_id, A.order_name, B.product, A.created_at, B.price_per_unit, B.quantity, B.id from orders A, order_items B where A.id=B.order_id) C "
	queryString3 := "left join (select order_item_id, sum(delivered_quantity) As delivered_quantity from deliveries group by order_item_id) D on C.id=D.order_item_id"
	queryString := queryString1 + queryString2 + queryString3

	orderString := " order by C.order_name"
	if search != "" {
		searchString := " where UPPER(c.order_name) LIKE Upper('%" + search + "%') or UPPER(C.product) LIKE Upper('%" + search + "%')"
		queryString += searchString
	}

	if startDate != "" {
		startDateString := "CAST(C.created_at AS date) >= CAST('" + startDate + "' AS date) AT TIME ZONE 'UTC'"
		endDateString := "CAST(C.created_at AS date) <= CAST('" + endDate + "' AS date) AT TIME ZONE 'UTC'"
		if search == "" && endDate == "" {
			queryString += " where " + startDateString
		} else if search != "" && endDate == "" {
			queryString += " AND " + startDateString
		} else if search == "" && endDate != "" {
			queryString += " where " + startDateString + " AND " + endDateString
		} else {
			queryString += " AND " + startDateString + " AND " + endDateString
		}
	} else {
		endDateString := "CAST(C.created_at AS date) <= CAST('" + endDate + "' AS date) AT TIME ZONE 'UTC'"
		if search == "" && endDate != "" {
			queryString += " where " + endDateString
		} else if search != "" && endDate != "" {
			queryString += " AND " + endDateString
		}
	}

	queryString += orderString

	// get data from sql database for order info
	orders := []Order{}

	data, err := postgreDB.Query(queryString)
	checkErr(err)

	for data.Next() {
		var customer_id string
		var order_name string
		var product string
		var created_at string
		var deliveried_amount float32
		var total_amount float32
		err = data.Scan(&customer_id, &order_name, &product, &created_at, &deliveried_amount, &total_amount)

		// get data from mongoDB database for customer info
		var customer bson.M
		if err = customersCollection.FindOne(
			ctx,
			bson.M{"user_id": customer_id},
		).Decode(&customer); err != nil {
			log.Fatal(err)
		}

		var company bson.M
		if err = cus_CompaniesCollection.FindOne(
			ctx,
			bson.M{"company_id": customer["company_id"]},
		).Decode(&company); err != nil {
			log.Fatal(err)
		}

		var customer_name = customer["name"].(string)
		var customer_company = company["company_name"].(string)

		// generate order object and add to data array
		var order Order
		order.OrderName = order_name
		order.CustomerName = customer_name
		order.CustomerCompany = customer_company
		order.ProductName = product
		order.OrderDate = created_at
		order.DeliveredAmount = deliveried_amount
		order.TotalAmount = total_amount

		orders = append(orders, order)

	}

	return orders

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
