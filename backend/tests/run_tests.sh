#!/bin/bash

# User tests

# generate random string as username
random_string=$(openssl rand -hex 8)
hurl -v --test --variable random_string=$random_string users.hurl

