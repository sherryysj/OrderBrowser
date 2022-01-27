-- for count data number
select count(*)
from orders A 
inner join order_items B on A.id=B.order_id 
inner join deliveries C on B.id=C.order_item_id
-- select according to user search input
where UPPER(A.order_name) LIKE Upper('%box%') or UPPER(B.product) LIKE Upper('%box%') 
-- select according to user start data input
and CAST(A.created_at AS date) > CAST('2020-1-5' AS date)
-- select according to user end data input
and CAST(A.created_at AS date) < CAST('2021-1-6' AS date);

-- for get data info
select A.customer_id, A.order_name, B.product, A.created_at, C.delivered_quantity*B.PRICE_PER_UNIT AS Deliveried_Amount, B.price_per_unit*B.quantity AS Total_Amount 
from orders A 
inner join order_items B on A.id=B.order_id 
inner join deliveries C on B.id=C.order_item_id
where UPPER(A.order_name) LIKE Upper('%box%') or UPPER(B.product) LIKE Upper('%box%') 
and CAST(A.created_at AS date) > CAST('2020-1-5' AS date)
and CAST(A.created_at AS date) < CAST('2021-1-6' AS date)