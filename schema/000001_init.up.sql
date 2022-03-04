CREATE TABLE users
(
    id      serial       NOT NULL unique,
    name    varchar(255) NOT NULL,
    balance int
);