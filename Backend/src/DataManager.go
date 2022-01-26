package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {

	// check arg input, if miss argument, do nothing
	if len(os.Args) < 3 {
		return
	}

	// arg1: Database - Mongo or Postgres
	// arg2: collection or table name (the letters at the end of data file names such as Orders for file "Test task - Mongo - Orders.csv")
	database := os.Args[1]
	name := os.Args[2]

	if os.Args[1] == "Postgres" {
		setPostgreSqlDB(database, name)
	} else {
		// for mongoDB, not implemented
	}

}

func setPostgreSqlDB(database string, name string) {

	// read database connect credentials
	config, err := os.ReadFile("../config/postgresql.config")
	checkErr(err)

	// connect database
	db, err := sql.Open("postgres", string(config))
	checkErr(err)

	// read data from csv file
	dataFile := "../data/Test task - Postgres - " + name + ".csv"
	data := readCsvData(dataFile)

	// set titles string and value string
	titles := data[0]
	titleString := ""
	valueString := ""
	for i, title := range titles {
		if i < len(titles)-1 {
			titleString += title + ","
			valueString += "$" + strconv.Itoa(i+1) + ","
		} else {
			titleString += title
			valueString += "$" + strconv.Itoa(i+1)
		}
	}

	// insert data into database
	for i, row := range data {

		queryString := "INSERT INTO " + name + " (" + titleString + ") VALUES (" + valueString + ")"

		if i > 0 {
			stmt, err := db.Prepare(queryString)
			checkErr(err)

			itemsAmount := len(titles)

			// below code need to improved
			switch itemsAmount {
			case 1:
				res, err := stmt.Exec(row[0])
				checkErr(err)
				fmt.Println(res)
			case 2:
				res, err := stmt.Exec(row[0], row[1])
				checkErr(err)
				fmt.Println(res)
			case 3:
				res, err := stmt.Exec(row[0], row[1], row[2])
				checkErr(err)
				fmt.Println(res)
			case 4:
				res, err := stmt.Exec(row[0], row[1], row[2], row[3])
				checkErr(err)
				fmt.Println(res)
			case 5:
				res, err := stmt.Exec(row[0], row[1], row[2], row[3], row[4])
				checkErr(err)
				fmt.Println(res)
			case 6:
				res, err := stmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5])
				checkErr(err)
				fmt.Println(res)
			case 7:
				res, err := stmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6])
				checkErr(err)
				fmt.Println(res)
			}

		}
	}

	db.Close()

}

// read csv data
func readCsvData(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
