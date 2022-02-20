# Order Browser
Web Application based on Vue.js and Golang

## Project setup steps

1. follow below Database setup instructions to set database config and insert data
2. direct to Backend folder and follow below Backend setup instructions to run the server
3. direct to Frontend folder and follow below Frontend setup instructions to run front end web application

## Database setup

1. create a folder called "config" under Backend folder to put database connection credentials
2. generate a file called "mongodb.config" under config folder
3. generate a file called "postgresql.config" under config folder
4. create a folder called "data" under Backend folder and put data csv files in it

### MongoDB

1. set a MongoDB database in Atlas (https://www.mongodb.com/atlas) or host a local database through MongoDB Compass
2. get the link credential of your databse from Atlas dashboard and copy it in mongodb.config file
3. direct to Backend/src folder in command line and follow below instructions to insert data into MongoDB

```
call python file in command line by "python DataManager.py start #arg1 #arg2"
#arg1: Database - "Mongo" or "Postgres"
#arg2: collection or table name (the letters at the end of data file names such as "Orders" for file "Test task - Mongo - Orders.csv")
command example: python DataManager.py start Mongo Orders
```

### PostgreSQL

1. host a PostgresSQL database through pgAdmin (https://www.pgadmin.org/)
2. direct to Backend/postgreSQL folder and copy scripts in DatabaseGenerator.sql file to paAdmin script board to generate tables
3. put following data into postgresql.config
   fileuser=yourUserName password=yourPassword dbname=yourDatabaseName sslmode=disable
4. direct to Backend/src folder in command line and follow below instructions to insert data into PostgreSQL

```
call golang file in command line by "go run DataManager.go #arg1 #arg2"
#arg1: Database - "Mongo" or "Postgres"
#arg2: collection or table name (the letters at the end of data file names such as "orders" for file "Test task - Postgres - orders.csv")
command example: go run DataManager.py Postgres orders
because of foreign key, sql data must insert in following order "1. orders, 2. order_items, 3. deliveries"
```

## Backend setup

run below code in command line at Backend folder

### Running server

```
go run server.go
```

## Frontend setup

run below code in command line at Fackend folder

### Install App

for first time running this project on your computer

```
npm install
```

### Run frontend app server

```
npm run serve
```

## What can be improved

1. I use a function to auto-insert data from different files, so it is hard to configue which data are numbers and which are stings,
   this causes data inserted into MongoDB are all strings, I am still trying to find a way to convert data to proper type without
   checking the each data file and write different code to convert data type for them
   
2. PostgresSQL data insert code (Line 74 to 104 in Backend/src/DataManager.go) should be improved to make it more efficiency

3. SQL query string (Line 102 to 132 in Backend/server.go) could be put in a separate file and read in when needed

4. Retrive data method in Backend/server.go could put in a seperate dababase manager file for a better quality of code management

5. A loading data message can be add to the page when data is loading

6. Orders per page auto reset to 5 every time using filter to retrive new data, it should be keep what user choose if user does not reload the App
