CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(128) UNIQUE,
    password VARCHAR(54),
    create_time TIMESTAMP
);