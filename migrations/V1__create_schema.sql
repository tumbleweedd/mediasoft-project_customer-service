CREATE TABLE if not exists offices
(
    uuid      uuid primary key,
    name      varchar(50)  not null,
    address   varchar(100) not null,
    created_at timestamp default now()
);

create table if not exists users
(
    uuid        uuid primary key,
    name        varchar(50) not null,
    office_uuid uuid        not null
        constraint fk_offices_users
            references offices (uuid),
    office_name varchar(20) not null,
    created_at  timestamp default now()
);

create table if not exists orders
(
    uuid uuid primary key,
    user_uuid  uuid not null
        constraint fk_users_orders
            references users (uuid)
);

create table if not exists order_items
(
    id serial primary key ,
    count int not null ,
    product_uuid uuid not null ,
    order_uuid uuid not null
        constraint fk_order_items
            references orders(uuid)
);