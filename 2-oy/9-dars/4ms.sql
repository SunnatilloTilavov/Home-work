
drop function genderuser;

select * from genderuser();

CREATE OR REPLACE FUNCTION genderuser(gender varchar)
RETURN INTEGER
LANGUAGE plpgsql
AS
$$
BEGIN
  RETURN(
  SELECT count(user_gender)
  FROM users WHERE user_gender=gender);
END;
$$;






CREATE OR REPLACE FUNCTION genderuser()
RETURNS TABLE (user_gender VARCHAR, count1 INT)
LANGUAGE plpgsql
AS
$$
BEGIN
  RETURN QUERY
  SELECT user_gender, COUNT(user_gender) AS count1
  FROM users
  GROUP BY user_gender;
  RETURN;
END;
$$;
