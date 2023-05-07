drop table if exists customer.users;

create table if not exists customer.users
(
    uuid        uuid primary key,
    name        varchar(50) not null,
    office_uuid uuid        not null
        constraint fk_offices_users
            references offices (uuid),
    office_name varchar(20) not null,
    created_at  timestamp default now()
);

