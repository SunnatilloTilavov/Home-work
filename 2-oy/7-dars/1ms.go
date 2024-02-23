1) select Person.firstName,Person.lastName,Address.city,Address.state 
from Person left join Address
ON Person.personId=Address.personId;

2)select distinct Customers.name AS Customers 
from Customers  left join Orders 
on Customers.id=Orders.customerId where Orders.customerId is null;

3)SELECT *
FROM (
    SELECT activity_date AS day, COUNT(DISTINCT user_id) AS active_users
    FROM Activity
    GROUP BY activity_date
) AS subquery
WHERE active_users > 1;

4)select EmployeeUNI.unique_id,Employees.name 
from Employees left join EmployeeUNI 
on Employees.id=EmployeeUNI.id ;

5)SELECT 
Product_id 
FROM Products where low_fats=recyclable and low_fats="Y" ;

6)SELECT event_day AS day, emp_id ,
sum(out_time-in_time) As total_time
FROM Employees
group by  event_day,emp_id;

7)select date_id,make_name,
count(distinct(lead_id)) as unique_leads,
count(distinct(partner_id)) as unique_partners from 
DailySales Group by date_id,make_name;

8)select teacher_id,count(distinct(subject_id)) as cnt
 from Teacher group by teacher_id;

 9)SELECT tweet_id FROM Tweets
 WHERE LENGTH(content) > 15;

 10)select Users.name as NAME,sum(Transactions.amount) AS BALANCE 
 from Users join Transactions  on Users.account=Transactions.account
  group by Transactions.account limit 1;

 11)select product.product_name, sales.year, sales.price 
 from product join sales 
 on product.product_id = sales.product_id

 12)select distinct author_id as id from Views where author_id=viewer_id 
 order by author_id;

 13)select employee.name, bonus.bonus
 from employee  left join bonus  on employee.empid = bonus.empid
 where bonus.bonus <1000 or bonus is null;

 14)select * FROM cinema WHERE (id % 2 != 0)
 AND (description != "boring") ORDER BY rating DESC;
 
 15)select actor_id,director_id from ActorDirector
 group by actor_id,director_id having count(timestamp)>=3 