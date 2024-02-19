sudo -u postgres psql -d sqldatabase
\l
\c sqldatabase 
Create table Student(id Serial,name varchar(34),age int,country varchar(34),salary int)
\d student;
Insert into Student(name,age,country,salary) Values('abs',12,'tashkent',1234),
('qwe',15,'asdefr',321),
('pqw',22,'ihmhmy',168),
('qad ',32,'oioqwe',574),
('gabs',12,'wqtqashkent',345),
('aaa',31,'asdtashkent',4321),
('adawj',42,'tashkent',423);
select *from Student;
select MIN(age) from Student;
select MAX(salary) from Student;
select COUNT(id) from Student;
select AVG(age) from Student;
select SUM(salary) from Student;