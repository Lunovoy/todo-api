CREATE TABLE users
(
    id serial primary key,
    name varchar(255) not null,
    login varchar(255) not null unique,
    password_hash varchar(255) not null
);