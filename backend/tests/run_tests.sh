#!/bin/bash

username=$(openssl rand -hex 8)
password=$(openssl rand -hex 8)

echo "Running user tests..."
hurl --test \
    --variable username=$username \
    --variable password=$password \
    users.hurl

item_name=$(openssl rand -hex 8)
item_units=2.2

echo "Running item tests..."
hurl --test \
    --variable username=$username \
    --variable password=$password \
    --variable item_name=$item_name \
    --variable item_units=$item_units \
    items.hurl

