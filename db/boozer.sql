-- NOTES: serial = auto incrementing number
CREATE TABLE users(
    user_id     SERIAL          PRIMARY KEY,
    username    VARCHAR(20)     NOT NULL,
    password    VARCHAR(256)    NOT NULL,
    created     TIMESTAMP       NOT NULL
);

CREATE TABLE items(
    item_id     SERIAL          PRIMARY KEY,
    name        VARCHAR(40)     NOT NULL,
    units       FLOAT           NOT NULL,
    added       TIMESTAMP       NOT NULL
);

CREATE TABLE consumptions(
    consumption_id  SERIAL      PRIMARY KEY,
    item_id         SERIAL      REFERENCES items,
    user_id         SERIAL      REFERENCES users,
    time            TIMESTAMP   NOT NULL
);

