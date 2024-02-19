Create table brands(
    id uuid DEFAULT gen_random_uuid() Primary key,
    year date,
    name varchar(25),
    stock boolean DEFAULT false,
    rate int);

Create table products(
    id uuid DEFAULT gen_random_uuid(),
    name varchar(25),
    barand_id uuid References  brands(id),
    price int,
    country varchar
);

Insert into  brands(year,name,stock,rate)
 values('2023-12-12','asus',true,5),
('2023-12-12','del',false,5),
('2023-12-12','hp',true,3);
Insert into products(name,barand_id,price,country) 
values('qwasd','669ba3b7-6037-48d1-8a0a-d88f6f24ceff',123123,'tashkent'),
('addwasd','3da77dc8-bb3f-4d1f-bf26-3b1c483c0c8f',55555,'ssss'),
('ssssss','2ab914be-c4ff-41bb-851f-0662aa9960b7',1111,'aaaa'),
('dddddd','2ab914be-c4ff-41bb-851f-0662aa9960b7',1113241,'sgs'),
('123124134','669ba3b7-6037-48d1-8a0a-d88f6f24ceff',234234,'aekfhkea');
SELECT *FROM brands where like name(%aa);
SELECT *FROM products;

SELECT *FROM brands order by name asc /*(asc or desc)*/
SELECT *FROM products where name IN('name','name')
SELECT *FROM products where population between 10000 and 300000

alter table brands add column creat_att_name time;
