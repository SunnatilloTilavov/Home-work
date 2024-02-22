sudo -u postgres psql -d sqldatabase
\l
\c sqldatabase 
\dt
CREATE TABLE Country(
    countryid uuid DEFAULT gen_random_uuid() Primary key,
    country_name varchar(25)
);


CREATE TABLE City(
    City_id uuid DEFAULT gen_random_uuid() Primary key,
    City_name varchar(25),
    country_id uuid References Contry(countryid)
);
INSERT INTO Country(country_name) values('Uzbekistan'),
('USA'),
('Kanada'),
('China'),
('Germany');
INSERT INTO City (City_name, country_id) VALUES
('Tashkent', '485f52cf-0afc-4400-9ead-29a1a77061ae'),
('Samarqand', '485f52cf-0afc-4400-9ead-29a1a77061ae'),
('New-York', '933b50b2-e3bf-4c3f-808e-4c0e0daf087b'),
('Chicago', '933b50b2-e3bf-4c3f-808e-4c0e0daf087b'),
('Toronto', '66502cff-6a60-428a-8346-2ed989751242'),
('Ottawa', '66502cff-6a60-428a-8346-2ed989751242'),
('Shanghai', '0e89287a-f98c-4a2a-bf4a-273015210429'),
('Guangzhou', '0e89287a-f98c-4a2a-bf4a-273015210429'),
('Berlin', 'd7980471-0d5f-44a8-9b6f-c3d998cd3da4'),
('Hamburg', 'd7980471-0d5f-44a8-9b6f-c3d998cd3da4');

SELECT City.City_name, Country.country_name FROM City JOIN Country ON City.country_id = Country.countryid;
SELECT AVG(length(City_name)) From city;
SELECT COUNT(City_name) FROM City;
//SELECT Country.country_name from Country JOIN CITY  WHERE (ON COUNT(City.country_id = Country.countryid)>3);
SELECT *FROM CITY WHERE City_name like'A%' ;

ALTER TABLE City ADD COLUMN population INT;

UPDATE City
SET population = CASE
    WHEN City_name = 'Tashkent' THEN 2000000
    WHEN City_name = 'Samarqand' THEN 700000
    WHEN City_name = 'New-York' THEN 8500000
    WHEN City_name = 'Chicago' THEN 2700000
    WHEN City_name = 'Toronto' THEN 2800000
    WHEN City_name = 'Ottawa' THEN 1000000
    WHEN City_name = 'Shanghai' THEN 24000000
    WHEN City_name = 'Guangzhou' THEN 14000000
    WHEN City_name = 'Berlin' THEN 3700000
    WHEN City_name = 'Hamburg' THEN 1800000
END;

SELECT SUM(population) FROM CITY;
UPDATE City  SET population=3000000 WHERE City_name='Tashkent';
DELETE FROM City WHERE population<100000;

/// TIME ADD
ALTER TABLE City ADD COLUMN TIME_ADD TIME;
UPDATE CITY SET City_name='TASHKENT1',TIME_ADD=current_timestamp WHERE City_id_id='6ba1eaa2-9335-4ce7-8661-3695d476d5ba';
UPDATE CITY SET population=55555555,TIME_ADD=current_timestamp WHERE City_name='Toronto';

SELECT City_name, COUNT(City_name) AS Soni
FROM City
GROUP BY City_name
HAVING COUNT(City_name) > 2;

