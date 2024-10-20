## Running
With docker & docker compose installed, run `docker compose up -d`
> **Note**: Remove the `-d` to keep the container running in the foreground.

### Shutting down
`docker compose down`

## Connecting
1. Run `docker ps` and find the ID of the container.
2. Run `docker exec -it <ID> psql -U postgres`
