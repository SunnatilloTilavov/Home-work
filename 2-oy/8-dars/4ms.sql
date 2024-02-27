Create Table order_products (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(20),
    price integer,
    order_id uuid references orders(id)
)
 
Create Table orders (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id integer,
    amount integer,
    count int
)


Insert Into orders (external_id, amount,count) Values(456789,100000),
(123123,120000,1),
(182131,21333,1),
(123123,2749713,1),
(24234,139595,1);


Insert Into order_products (name, price, order_id) 
 VALUES('Apple',1000,'42c730c6-ee71-4902-8867-500ab305f62a'),
 ('samsung',2000,'cacb8cec-2992-438f-97d9-5d4da5e796b7'),
 ('mi',3000,'8e6a0f13-07bc-4643-8abf-fcf1e2fe380f'),
 ('vivo',1232,'0f6234e2-3ce0-4902-8a71-d3ffae623aa1'),
 ('awdaw',1221,'59d20fb2-2825-4d92-a87c-ee3694883e1d');