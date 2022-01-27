-- Database: OrderBrowser

/* 
-- for reset use
delete from deliveries;
delete from order_items;
delete from orders;

drop table Deliveries;
drop table Order_items;
drop table Orders;
*/

CREATE TABLE Orders(
    ID varchar(64) NOT NULL,
	created_at varchar,
	order_name varchar(255),
	customer_id varchar(64),
	PRIMARY KEY (ID)
);

CREATE TABLE Order_items(
    ID varchar(64) NOT NULL,
	order_id varchar(64),
	price_per_unit DECIMAL,
	quantity int,
	product varchar(255),
	PRIMARY KEY (ID),
	FOREIGN KEY (order_id) REFERENCES Orders(ID)
);

CREATE TABLE Deliveries(
    ID varchar(64) NOT NULL,
	order_item_id varchar(64),
	delivered_quantity int,
	PRIMARY KEY (ID),
	FOREIGN KEY (order_item_id) REFERENCES Order_items(ID)
);




	