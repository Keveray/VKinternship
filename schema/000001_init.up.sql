CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null

);

CREATE TABLE VKinternship_lists
(
    id serial not null unique,
    title varchar(255) not null,
    descriptions varchar(255)
);

CREATE TABLE user_lists
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    list_id int references VKinternship_lists(id) on delete cascade not null
);

CREATE TABLE vk_lenta
(
    id serial not null unique,
    title varchar(255) not null,
    descriptions varchar(255),
    viewed boolean not null default false
);

CREATE TABLE lists_lenta
(
    id serial not null unique,
    lenta_id int references vk_lenta(id) on delete cascade not null,
    list_id int references VKinternship_lists(id) on delete cascade not null
);
