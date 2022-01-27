import pymongo
import sys
import csv

# arg1: Database - Mongo or Postgres
# arg2: collection or table name (the letters at the end of data file names such as Orders for file "Test task - Mongo - Orders.csv")
def start(database, name):
    if database == "Mongo":
        mongoDB(name)
    else:
        # not set database yet
        postgres(name)  

def mongoDB(name):
    
    # read data from data file
    filePath = "../data/Test task - Mongo - " + name +".csv"
    data = list(csv.reader(open(filePath, mode='r')))

    # connect database
    dataFile = open("../config/mongodb.config", "r")
    client = pymongo.MongoClient(dataFile.readline())
    database = client["OrderBrowser"]
    collection = database[name]

    # insert data into database
    for index in range(1,len(data)):
        newData = {}
        i = 0
        for item in data[index]:
            newData[data[0][i]] = item
            i += 1
        insertData = collection.insert_one(newData)

if __name__ == "__main__":
    args = sys.argv
    # args[0] = current file
    # args[1] = function name
    # args[2:] = function args : (*unpacked)
    globals()[args[1]](*args[2:])

