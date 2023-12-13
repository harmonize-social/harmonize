CREATE TABLE IF NOT EXISTS posts(
    id uuid PRIMARY KEY,
    user_id uuid references users (id),
    caption varchar(256) NOT NULL,
    type varchar(16),
    type_specific_id uuid
);

CREATE TABLE IF NOT EXISTS likes(
    id uuid PRIMARY KEY,
    post_id uuid references posts (id),
    user_id uuid references users (id)
);

CREATE TABLE IF NOT EXISTS comment(
    id uuid PRIMARY KEY,
    post_id uuid references posts (id),
    user_id uuid references users (id),
    reply_to_id uuid references comment (id),
    message varchar(256)
);

CREATE TABLE IF NOT EXISTS saved_posts(
    id uuid PRIMARY KEY,
    user_id uuid references posts (id),
    post_id uuid references posts (id)
);
