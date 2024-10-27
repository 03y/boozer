## Building
```
go get boozer
go build
```

## Running
Set environment variable for postgres connection:
```
DATABASE_URL='postgres://postgres:postgres@localhost:5001/boozer'
```
(or appropriate values set in `db/docker-compose.yml`.

Specify URL:port on running like:
```
./boozer localhost:6000
```

