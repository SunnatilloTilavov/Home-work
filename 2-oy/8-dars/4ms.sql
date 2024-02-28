Create Table order_products1 (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(20),
    price integer,
    order_id uuid references orders1(id),
    totalsum int
)
 
Create Table orders1 (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id integer,
    amount integer
)


Insert Into orders1 (external_id, amount) Values(456789,100000),
(123123,12000),
(182131,21333),
(123123,2749713),
(24234,139595);


Insert Into order_products1 (name, price, order_id) 
 VALUES('Apple',1000,'84678db0-c6c0-4360-8356-6db1e7b3f315'),
 ('samsung',2000,'06eaac23-a6b5-4824-9c3a-2c21f01b7b12'),
 ('mi',3000,'a8db2ae7-d655-41c9-9c71-a31b7786a3e3'),
 ('vivo',1232,'555ab086-13bf-41a4-9ab2-8bfb4388b81f'),
 ('awdaw',1221,'54b24040-1b2e-4392-9660-06f837fa44d5');

 Create or replace Function totalsum1();
 returns trigger
 language plpgsql
 as $$
 begin
 update orders1 set order_products1.totalsum=order_products1.price*orders1.amount from orders1 join order_product1 on order_product1.order_id=orders1.id  where id=NEW.order_id;
 return NEW;
 end;
 $$;

 Create trigger inc_totalsum1
 AFTER INSERT
 ON order_product1
 FOR EACH ROW
 EXECUTE PROCEDURE  totalsum1();
 
 select *from orders1;