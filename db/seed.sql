CREATE DATABASE boozer;

\c boozer

CREATE TABLE users(
    user_id     SERIAL          PRIMARY KEY,
    username    VARCHAR(20)     NOT NULL    UNIQUE,
    password    VARCHAR         NOT NULL,
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

INSERT INTO items (name, units, added) VALUES ('Früh Kölsch',2.4, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Rothaus Pils',2.6, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Ishii Orehi Pale Ale',2.4, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Exmoor Phoenix',2.8, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Moniack Sweet Mead',1.2, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Orkney Dark Island Pint',2.6, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Urban South Who Dat',2.8, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Hacker Pschorr Münchner Gold',2.8, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Augustinerbräu München Dunkel',2.8, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Central City Racer IPA',2.3, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Townsend Dinner Ale',2.4, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Mòr Session IPA Pint',2.1, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Harviestown Bitter and Twisted Pint',2.2, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Belhaven Best Pint',1.8, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Williams Bros Craft Lager',1.8, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Williams Bros El Perro Negro',2.0, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Brooklyn Pilsner',2.0, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Margiotta Lager',1.8, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Bud Light Pint',1.9, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Worthington''s Creamflow',1.9, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Grainstone Rutland Beast',3.1, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Craingorm Gold Bottle',2.3, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Craingorm Trade Winds Bottle',2.2, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Old Speckled Hen',2.6, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Krombacher Pils',3.2, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Fierce Fancy Juice',2.2, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Goose IPA Bottle',1.9, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Stella Pint',2.6, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Juicy Joker',2.2, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Tennents 500ml Can',2.0, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Williams Bros Birds and Bees',2.2, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Harviestown The Ridge Bottle',2.5, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Hopo',3.1, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Schöfferhofer Grapefruit',1.3, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Innis and Gunn 440ml Can',2.0, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Skye Black Bottle',2.3, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Skye Red Bottle',2.1, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Badger Hopping Hare',2.0, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Skye Blaven Bottle',2.4, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('West 4 Pint',2.3, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Guiness Pint',2.4, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Cuillin Brewery Pilsner',2.8, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Skye Golden Ale Pint',2.4, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Carling Pint',2.3, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Haviestoun  Schiehallion Bottle',2.4, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Tennents Pint',2.3, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Coop Golden Ale',2.0, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Stella 660ml bottle',3.0, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Inveralmond Ossian Bottle',2.1, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Pilsner Urquell Bottle',2.2, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Belhaven 80 Shilling Pint',2.2, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Seven Brothers Throwaway IPA Can',2.2, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Orkney Cliff Edge IPA Pint',2.7, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Doom Bar Pint',2.4, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Joker IPA',2.5, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Sainsbury''s IPA',2.6, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Tennents 440ml Can',1.8, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Birra Moretti 600ml',3, 1686604802);

