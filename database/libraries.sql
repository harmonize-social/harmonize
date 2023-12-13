CREATE TABLE IF NOT EXISTS connections(
    id uuid primary key,
    user_id uuid references users (id),
    access_token varchar(64),
    refresh_token varchar(64),
    expiry timestamptz
);

CREATE TABLE IF NOT EXISTS libraries(
    id uuid primary key,
    platform_id varchar(64),
    connection_id uuid references connections (id)
);

CREATE TABLE IF NOT EXISTS images(
    id uuid primary key,
    path varchar(64)
);

CREATE TABLE IF NOT EXISTS platforms(
    id uuid primary key,
    name varchar(32),
    icon_id uuid references images (id)
);
