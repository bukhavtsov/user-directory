--Create user table
create table if not exists "users"
(
    id         serial      not null
        constraint user_pk
            primary key,
    first_name varchar(35) not null,
    last_name  varchar(35) not null,
    img        varchar(255)
);

alter table "users"
    owner to postgres;

create unique index user_id_uindex
    on "users" (id);

--Insert Users to db
insert into "users" (first_name, last_name, img)
values ('Jack', 'Bronson', 'assets/images/user_icon_1.png');

insert into "users" (first_name, last_name, img)
values ('Bob', 'Martin', 'assets/images/user_icon_2.png');

insert into "users" (first_name, last_name, img)
values ('Mike', 'Ford', 'assets/images/user_icon_3.png');

insert into "users" (first_name, last_name, img)
values ('Tommy', 'White', 'assets/images/user_icon_4.png');

insert into "users" (first_name, last_name, img)
values ('Roman', 'Sim', 'assets/images/user_icon_5.png');

--Select all Users from db
select * from "users"