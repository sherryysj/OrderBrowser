# orderbrowser

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

### PostSQL

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

### running server

```
go run server.go
```

## Frontend

### Install (for first time running this project on your computer)

```
npm install
```

### Compiles and hot-reloads for development

```
npm run serve
```

### Drawbacks

1. I use a function to auto-insert data from different files, it is hard to configue which data are numbers and which are stings,
   this causes data inserted into MongoDB are all strings, I am still trying to find a way to convert data to proper type without
   checking the each data file and write different code to convert data type for them
2. PostgresSQL data insert code (Line 74 to 104 in Backend/src/DataManager.go) should be improved to make it more efficiency