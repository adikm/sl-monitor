CREATE TABLE users
(
    id       INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name     TEXT    NOT NULL,
    email    TEXT    NOT NULL UNIQUE,
    password TEXT    NOT NULL
);