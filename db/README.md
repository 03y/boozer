## Running
With docker & docker compose installed, run `docker compose up --build`.
> **Note**: Add `-d` to run the container in the background.

### Shutting down
`docker compose down`

## Connecting
1. Run `docker ps` and find the ID of the container.
2. Run `docker exec -it <ID> psql -U postgres`.
3. In here connect to the database with `\c boozer`.
4. Run SQL commands from here.
