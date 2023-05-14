create table if not exists customer.orders
(
    uuid uuid primary key,
    user_uuid  uuid not null
        constraint fk_users_orders
            references customer.users (uuid)
);