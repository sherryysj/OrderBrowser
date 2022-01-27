-- for get data info
-- join selected orders and deliveries below and show the data needed for from end 
select C.customer_id, C.order_name, C.product, C.created_at, D.delivered_quantity*C.PRICE_PER_UNIT AS Deliveried_Amount, C.price_per_unit*C.quantity AS Total_Amount 
-- join orders and order_items table to display orders data according to their items 
from (select A.customer_id, A.order_name, B.product, A.created_at, B.price_per_unit, B.quantity, B.id from orders A, order_items B where A.id=B.order_id) C
-- sum delivered_quantity for the same item of an order and then join deliveries into orders above to add delivered_quantity data
left join (select order_item_id, sum(delivered_quantity) As delivered_quantity from deliveries group by order_item_id) D
on C.id=D.order_item_id
-- filter according to user search input
where UPPER(c.order_name) LIKE Upper('%box%') or UPPER(C.product) LIKE Upper('%box%') 
-- select according to user start data input
and CAST(C.created_at AS date) >= CAST('2020-1-5' AS date) AT TIME ZONE 'UTC'
-- select according to user end data input
and CAST(C.created_at AS date) <= CAST('2021-1-6' AS date) AT TIME ZONE 'UTC'
-- sort by order name for easy reading 
order by C.order_name

-- for count data number
select count(*)
from (select A.customer_id, A.order_name, B.product, A.created_at, B.price_per_unit, B.quantity, B.id from orders A, order_items B where A.id=B.order_id) C
left join (select order_item_id, sum(delivered_quantity) As delivered_quantity from deliveries group by order_item_id) D
on C.id=D.order_item_id
where UPPER(c.order_name) LIKE Upper('%box%') or UPPER(C.product) LIKE Upper('%box%') 
and CAST(C.created_at AS date) >= CAST('2020-1-5' AS date) AT TIME ZONE 'UTC'
and CAST(C.created_at AS date) <= CAST('2021-1-6' AS date) AT TIME ZONE 'UTC'
