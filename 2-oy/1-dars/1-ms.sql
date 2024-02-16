sudo -u postgres psql
\l
\c sqldatabase;
drop table students ; 
Create table Students(id Serial ,First_name varchar ,Last_name varchar,Age int,Gender varchar,Score int,Addresss varchar,Email varchar,Phone_number varchar,Date_of_birth date);
\d students;
Insert into Students(First_name,Last_name,Age,Gender,Score,Addresss,Email,Phone_number,Date_of_birth ) VAlues('A1','b1',11,'male',60,'Toshkent','agr@gmail.com','189382936123','2022-12-12'),
('A2','b2',12,'female',75,'Toshkent12','agr31@gmail.com','9189382936123','2023-12-12'),
('A3','b3',21,'male',55,'Toshkent4','12agr@gmail.com','189382936123','2021-10-1'),
('A4','b5',16,'female',40,'Toshkent12','ag213r@gmail.com','189382936123','2020-12-12'),
('A5','b5',12,'female',75,'Toshkent12','agr31@gmail.com','9189382936123','2023-12-12'),
('A6','b6',21,'male',55,'Toshkent4','12agr@gmail.com','189382936123','2021-10-1'),
('A7','b7',16,'female',40,'Toshkent12','ag213r@gmail.com','189382936123','2020-12-12');
select *from Students;
update students SET  First_name='A8' where id=2;
update students SET  Gender='female' where Gender='male';
update students SET  Score=100 where Score>=100;
update students SET  Last_name='Rename',First_name='Rename2',Addresss='asdf',Phone_number='9237123d',Date_of_birth='1242-4-25' where id<=4;
delete from Students where Score<=60;



/*Auto use date*/
CREATE TABLE YTable (
    id Serial,
    other_column VARCHAR(255),
    timestamp_column TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
Insert into YTable(other_column) VAlues ('absd')