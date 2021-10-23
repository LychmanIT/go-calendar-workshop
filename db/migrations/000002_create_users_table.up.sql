CREATE TABLE users (
    id VARCHAR ( 100 ) primary key,
    username VARCHAR ( 50 ) UNIQUE NOT NULL,
    password VARCHAR ( 100 ) NOT NULL
);