1)select round(avg(POPULATION),0) from CITY;

2)Select max(months * salary), count(months * salary) from Employee Where (months * salary) = (Select max(months * salary) from Employee);

3)select round(sum(LAT_N),2),round(sum(LONG_W),2) from STATION;

4)

5)select COUNTRY.Continent,FLOOR(avg(CITY.Population)) from COUNTRY JOIN CITY on CITY.COUNTRYCODE=COUNTRY.CODE  GROUP BY  COUNTRY.Continent ;

6)