drop table if exists customer.offices;

CREATE TABLE if not exists customer.offices
(
    uuid      uuid primary key,
    name      varchar(50)  not null,
    address   varchar(100) not null,
/*    company_uuid uuid          not null
        constraint fk_companies_offices
            references companies (uuid),*/
    created_at timestamp default now()
)