# DB

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

## Schema
Three tables: users, items and uses.

### Users
| Field         | Datatype      | Key | Comments  |
| ------------- |:-------------:|:---:|:---------:|
| user\_id      | int(20)       | PK  | AUTO-INC  |
| username      | varchar(20)   |     |           |
| password      | varchar       |     | argon2    |
| created       | int           |     | unix time |

### Items
| Field         | Datatype      | Key | Comments  |
| ------------- |:-------------:|:---:|:---------:|
| item\_id      | int(20)       | PK  | AUTO-INC  |
| name          | varchar(40)   |     |           |
| units         | float         |     |           |
| added         | int           |     | unix time |

### Consumptions
| Field         | Datatype      | Key | Comments  |
| ------------- |:-------------:|:---:|:---------:|
| consumption\_id       | int(20)       | PK  | AUTO-INC  |
| item\_id      | int(20)       | FK  |           |
| user\_id      | int(20)       | FK  |           |
| time          | int           |     | unix time |
