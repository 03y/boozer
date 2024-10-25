# syntax=docker/dockerfile:1
FROM postgres:latest
COPY boozer.sql /docker-entrypoint-initdb.d/
COPY users.sql /docker-entrypoint-initdb.d/
COPY items.sql /docker-entrypoint-initdb.d/
COPY consumptions.sql /docker-entrypoint-initdb.d/
