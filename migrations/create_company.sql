drop table if exists customer.companies;

create table if not exists customer.companies
(
    id   serial primary key,
    name varchar(50) not null,
    uuid uuid        not null unique
)