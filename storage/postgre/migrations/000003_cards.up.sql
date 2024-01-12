CREATE TABLE IF NOT EXISTS cards (
    id bigserial PRIMARY KEY,
    user_id text NOT NULL,
    money double precision
);