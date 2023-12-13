CREATE TABLE IF NOT EXISTS users (
    id uuid primary key,
    email varchar(64) unique not null,
    username varchar(64) unique not null,
    password_hash varchar(256) not null
);

CREATE TABLE IF NOT EXISTS sessions (
    id uuid primary key,
    user_id uuid references users (id) not null,
    expiry timestamptz not null
);

CREATE TABLE IF NOT EXISTS subscriptions (
    user_id uuid primary key references users (id) not null,
    stripe_subscription varchar(128) not null
);

CREATE TABLE IF NOT EXISTS follows(
    id uuid PRIMARY KEY,
    followed_id uuid references users (id),
    follower_id uuid references users (id),
    date timestamptz
);
