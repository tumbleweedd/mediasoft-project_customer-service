create table if not exists customer.order_items
(
    id serial primary key ,
    count int not null ,
    product_uuid uuid not null ,
    order_uuid uuid not null
        constraint fk_order_items
            references customer.orders(uuid)
);