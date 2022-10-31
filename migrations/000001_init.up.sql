CREATE TABLE users
(
    id serial primary key,
    name varchar(255) not null,
    login varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE lists
(
    id serial primary key,
    title varchar(255) not null
);

CREATE TABLE items
(
    id serial primary key,
    title varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE user_list
(
    id      serial primary key,
    user_id int references users (id) on delete cascade not null,
    list_id int references lists (id) on delete cascade not null
);

CREATE TABLE list_item
(
    id      serial primary key,
    item_id int references items (id) on delete cascade not null,
    list_id int references lists (id) on delete cascade not null
);
