CREATE DATABASE boozer;

\c boozer

CREATE TABLE users(
    user_id     SERIAL          PRIMARY KEY,
    username    VARCHAR(20)     NOT NULL    UNIQUE,
    password    VARCHAR         NOT NULL,
    created     INT             NOT NULL,
    recap_2025  JSONB
);

CREATE TABLE items(
    item_id     SERIAL          PRIMARY KEY,
    user_id     SERIAL          REFERENCES users,
    name        VARCHAR(40)     NOT NULL    UNIQUE,
    units       FLOAT           NOT NULL,
    added       INT             NOT NULL
);

CREATE TABLE consumptions(
    consumption_id  SERIAL      PRIMARY KEY,
    item_id         INT         REFERENCES items,
    user_id         SERIAL      REFERENCES users,
    time            INT         NOT NULL,
    price           FLOAT       NULL
);

CREATE TYPE reason_type AS ENUM ('name', 'units', 'duplicate', 'other');

CREATE TABLE item_reports(
    report_id   SERIAL          PRIMARY KEY,
    item_id     SERIAL          REFERENCES ITEMS,
    user_id     SERIAL          REFERENCES USERS,
    reason      reason_type     NOT NULL,
    created     INT             NOT NULL
);
