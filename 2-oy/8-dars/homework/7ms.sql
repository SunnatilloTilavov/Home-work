create or replace function gettotal(order_idd uuid)
   returns int
   language plpgsql
  as
$$
declare 
Totalprice int;
begin
select sum(order_products.price*orders.amount)
into totalprice
from order_products join orders on order_products.order_id=orders.id
where orders.id=order_idd;

return totalprice;
end;
$$;

select gettotal('42c730c6-ee71-4902-8867-500ab305f62a')