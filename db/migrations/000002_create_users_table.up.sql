CREATE TABLE users (
    id SERIAL primary key,
    username VARCHAR ( 50 ) UNIQUE NOT NULL,
    password VARCHAR ( 100 ) NOT NULL
);