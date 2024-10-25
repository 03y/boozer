CREATE DATABASE boozer;

\c boozer

CREATE TABLE users(
    user_id     SERIAL          PRIMARY KEY,
    username    VARCHAR(20)     NOT NULL    UNIQUE,
    password    VARCHAR(256)    NOT NULL,
    created     INT             NOT NULL
);

CREATE TABLE items(
    item_id     SERIAL          PRIMARY KEY,
    name        VARCHAR(40)     NOT NULL    UNIQUE,
    units       FLOAT           NOT NULL,
    added       INT             NOT NULL
);

CREATE TABLE consumptions(
    consumption_id  SERIAL      PRIMARY KEY,
    item_id         SERIAL      REFERENCES items,
    user_id         SERIAL      REFERENCES users,
    time            INT         NOT NULL
);

