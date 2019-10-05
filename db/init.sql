--Create user table
create table "user"
(
    id         serial  not null
        constraint user_pk
            primary key,
    first_name integer not null,
    last_name  integer not null,
    img_name   text,
    img        bytea
);

alter table "user"
    owner to postgres;

create unique index user_id_uindex
    on "user" (id);