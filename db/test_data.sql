-- create users
INSERT INTO users (username, password, created) VALUES ('Pibble', 'lemur', EXTRACT(epoch FROM now()));
INSERT INTO users (username, password, created) VALUES ('Glorp', 'lemur', EXTRACT(epoch FROM now()));
INSERT INTO users (username, password, created) VALUES ('Gleeb', 'lemur', EXTRACT(epoch FROM now()));
INSERT INTO users (username, password, created) VALUES ('Gnarp', 'lemur', EXTRACT(epoch FROM now()));

-- add some items
INSERT INTO items (name, units, added) VALUES ('Tennents', 2.272, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Staropramen', 3.3, EXTRACT(epoch FROM now()));
INSERT INTO items (name, units, added) VALUES ('Guinness', 2.3, EXTRACT(epoch FROM now()));

-- add some consumptions
-- todo: dont hardcode the user and item ids here
INSERT INTO consumptions (user_id, item_id, time) VALUES (1, 1, EXTRACT(epoch FROM now()));
INSERT INTO consumptions (user_id, item_id, time) VALUES (1, 2, EXTRACT(epoch FROM now()));
INSERT INTO consumptions (user_id, item_id, time) VALUES (1, 2, EXTRACT(epoch FROM now()));
INSERT INTO consumptions (user_id, item_id, time) VALUES (2, 3, EXTRACT(epoch FROM now()));
INSERT INTO consumptions (user_id, item_id, time) VALUES (2, 3, EXTRACT(epoch FROM now()));
INSERT INTO consumptions (user_id, item_id, time) VALUES (2, 3, EXTRACT(epoch FROM now()));
INSERT INTO consumptions (user_id, item_id, time) VALUES (3, 1, EXTRACT(epoch FROM now()));
INSERT INTO consumptions (user_id, item_id, time) VALUES (3, 3, EXTRACT(epoch FROM now()));
INSERT INTO consumptions (user_id, item_id, time) VALUES (4, 2, EXTRACT(epoch FROM now()));
