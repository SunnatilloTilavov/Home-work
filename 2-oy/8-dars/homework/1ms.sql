create or replace function increase_pr_count()
return trigger
LANGUAGE plpgsql
as $$
begin
update orders set products_count=products_count+1 where id= NEW.order_id;
return NEW;
end;
$$;


DROP trigger increase_order_producTS_count on order_products;


create trigger increase_order_producTS_count
after insert /*or update*/ on order_products
for each  row
execute procedure increase_pr_count();

select * from orders;