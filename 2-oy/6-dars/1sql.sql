sudo -u postgres psql -d sqldatabase
\c sqldatabase 

CREATE TABLE Customer(
    id uuid DEFAULT gen_random_uuid() Primary key,
    name varchar(25),
    phone varchar(25) UNIQUE,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp,
);

CREATE TABLE Adress(
    id uuid DEFAULT gen_random_uuid() ,
    adress varchar(25),
    customer_id uuid References Customer(id),

);

INSERT INTO Customer (name,phone) values('adasda',1234567),
('awdaw',9173231),
('awdwad',11289738),
('awdawdad',6213137);

INSERT INTO Adress(adress,customer_id) values('advw','ba6d5d20-a1e9-4ffe-a986-f6e4b3cc572b'),
('adva','ba6d5d20-a1e9-4ffe-a986-f6e4b3cc572b'),
('advada','ba6d5d20-a1e9-4ffe-a986-f6e4b3cc572b'),
('advw','1f6e0600-008b-494b-b54d-025ed2a4af0c'),
('asdasddvw','1f6e0600-008b-494b-b54d-025ed2a4af0c'),
('wadha','1f6e0600-008b-494b-b54d-025ed2a4af0c');
('adasda','a0a4d699-fa44-4ed0-8ab6-b21aca899823')

update customer SET deleted_at=NOW() where id='ba6d5d20-a1e9-4ffe-a986-f6e4b3cc572b';

INSERT INTO Customer (name,phone) values('adasda',1234567);

select Customer.name,Customer.phone,Adress.adress ,Customer.created_at from Customer left Join Adress ON Customer.id=Adress.customer_id;


