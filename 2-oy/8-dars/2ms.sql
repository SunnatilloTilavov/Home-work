create or replace function get_count_name(names varchar(50))
   returns int
   language plpgsql
  as
$$
declare 
namecount int;
begin
select count(*)
into namecount
from Customer1
where name=names;

return namecount;
end;
$$;

select get_count_name('awdaw')