CREATE TABLE notifications
(
    id        INTEGER   NOT NULL PRIMARY KEY AUTOINCREMENT,
    timestamp TIMESTAMP NOT NULL,
    user_id   INTEGER   NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users (id)
);
