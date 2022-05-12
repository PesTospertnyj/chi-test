create table IF NOT EXISTS books
(
    id     serial
        constraint books_pk
            primary key,
    author varchar(1024) not null,
    title  varchar(1024) not null
);

alter table books
    owner to api;

create unique index IF NOT EXISTS books_id_uindex
    on books (id);

