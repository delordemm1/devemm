CREATE TABLE users (
    id serial primary key,
    first_name text not null,
    last_name text not null,
    bio text not null,
    username text not null,
    password text not null,
    email text not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);