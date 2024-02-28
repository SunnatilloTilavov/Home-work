
CREATE TABLE users(user_id SERIAL PRIMARY KEY,
user_first_name VARCHAR(30),
user_last_name VARCHAR(30),
user_email_id VARCHAR(50),
user_gender VARCHAR(1),
user_unique_id VARCHAR(15),
user_phone_no VARCHAR(20),
user_dob DATE,
created_ts TIMESTAMP);


CREATE OR REPLACE FUNCTION createdate()
RETURNS INTEGER
LANGUAGE plpgsql
AS
$$
BEGIN
    RETURN(
    SELECT extract(YEAR FROM created_ts) AS year_count 
    FROM users 
    GROUP BY  year_count
    ORDER BY  count(*) DESC 
    LIMIT 1);
END;
$$;

drop function  createdate;

select * from createdate();