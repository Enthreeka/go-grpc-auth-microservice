CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create type role as enum ('admin','user');

create table if not exists "user"
(
    id uuid default uuid_generate_v4(),
    email varchar(100) not null unique,
    password varchar(200) not null,
    role role not null default 'user',
    primary key (id)
);

create table if not exists token
(
    id int generated always as identity,
    expires_at timestamp not null,
    created_at timestamp with time zone not null default now(),
    fingerprint varchar(200) not null,
    refresh_token uuid not null ,
    user_id uuid ,
    primary key (id),
    foreign key (user_id)
        references "user" (id) on delete cascade
);
