CREATE TABLE notifications
(
    id        SERIAL PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL,
    user_id   INTEGER   NOT NULL,
    weekdays  INT8      NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
