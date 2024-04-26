CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    surname text NOT NUll,
    email text Not null,
    password text not null
);