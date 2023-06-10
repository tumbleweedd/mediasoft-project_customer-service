-- SCHEMA
CREATE TABLE if not exists offices
(
    uuid       uuid primary key,
    name       varchar(50)  not null,
    address    varchar(100) not null,
    created_at timestamp default now()
);

-- SEED
insert into offices
    (uuid, name, address)
VALUES ('bfe1ff31-cef6-4e9a-acb7-bb38ec60b2c5', 'test office 1', 'test address 1');
insert into offices
(uuid, name, address)
VALUES ('6d0190d7-403f-4882-a7c9-2608f910731a', 'test office 2', 'test address 2');
insert into offices
(uuid, name, address)
VALUES ('36ff085b-3e49-4f37-ac71-ae4daf6ae0c9', 'test office 3', 'test address 3');

