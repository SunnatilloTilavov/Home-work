sudo -u postgres psql -d sqldatabase

CREATE TABLE Customer1(
    id uuid DEFAULT gen_random_uuid() Primary key,
    name varchar(25),
    phone varchar(25) UNIQUE,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp
);

CREATE UNIQUE INDEX index_phone
ON Customers(phone,deleted_at);

SELECT 
    phone, 
    indexname, 
    indexdef 
FROM 
    pg_indexes 
WHERE 
    tablename = 'Customer1';

INSERT INTO Customer1 (name,phone) values('adasda',1234567),
('awdaw',9173231),
('awdwad',11289738),
('awdawdad',6213137);


create or replace function getmaxid;
   returns int
   language plpgsql
  as
$$
declare 
namecount int;
begin
select count(*)
into namecount
from Customer1
where name=names;

return namecount;
end;
$$;

select getmaxid;