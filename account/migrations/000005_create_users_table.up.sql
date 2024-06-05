create extension if not exists citext;

CREATE TABLE IF NOT EXISTS users
(
    id            bigserial PRIMARY KEY,
    created_at    timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at    timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    fname         text                        NOT NULL,
    sname         text                        NOT NULL,
    email         citext UNIQUE               NOT NULL,
    password_hash bytea                       NOT NULL,
    activated     bool                        NOT NULL,
    user_role     varchar(50)                 not null,
    version       integer                     NOT NULL DEFAULT 1
);

