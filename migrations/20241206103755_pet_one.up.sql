CREATE TABLE users
(
    id            INTEGER     NOT NULL PRIMARY KEY,
    name          VARCHAR     NOT NULL,
    email         VARCHAR     NOT NULL UNIQUE,
    age           INTEGER     NOT NULL,
    password_hash VARCHAR(32) NOT NULL
);