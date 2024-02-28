CREATE OR REPLACE FUNCTION returnname(year1 INT)
RETURNS TABLE (first_name VARCHAR, last_name VARCHAR, created_date DATE)
LANGUAGE plpgsql
AS
$$
BEGIN
  RETURN QUERY
  SELECT user_first_name, user_last_name, created_ts::DATE
  FROM users
  WHERE EXTRACT(YEAR FROM created_ts) = year1;
  RETURN;
END;
$$;

drop function returnname;

select * from returnname(5);