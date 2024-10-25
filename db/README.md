## Running
With docker & docker compose installed, run `docker compose up -d --build`.
> **Note**: Remove the `-d` to keep the container running in the foreground.

### Shutting down
`docker compose down`

## Connecting
1. Run `docker ps` and find the ID of the container.
2. Run `docker exec -it boozer_test_db psql -U postgres`.
3. In here connect to the database with `\c boozer`.
