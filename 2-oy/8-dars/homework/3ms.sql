CREATE INDEX index_order_id
ON order_products(order_id);

Explain select *from order_products where order_id='42c730c6-ee71-4902-8867-500ab305f62a'