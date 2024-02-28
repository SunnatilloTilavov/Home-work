
drop function bornuser;
select * from bornuser(5);



CREATE OR REPLACE FUNCTION bornuser(abs int)
RETURNS TABLE (first_name VARCHAR, last_name VARCHAR, created_date DATE)
LANGUAGE plpgsql
AS
$$
BEGIN
  RETURN QUERY
  SELECT user_first_name, user_last_name, created_ts::DATE
  FROM users
  WHERE  extract(Month FROM created_ts)=abs;
  RETURN;
END;
$$;