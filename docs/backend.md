# Backend

- [Go](https://go.dev).
  - [Gin for HTTP API](https://pkg.go.dev/github.com/gin-gonic/gin).
  - [PGX for DB connection](https://github.com/jackc/pgx).

## Outline

Here is the requested information in markdown tables:

### Endpoints

| Method | Route                    | Parameters                                      | Notes                                            |
| ------ | ------------------------ | ----------------------------------------------- | ------------------------------------------------ |
| POST   | `/submit/item`           | `name` (string), `type` (string), `abv` (float) | requires JWT token (obtain from `/authenticate`) |
| POST   | `/submit/consumption`    | `item_id` (int), `amount` (int)                 | requires JWT token (obtain from `/authenticate`) |
| GET    | `/item/:item_id`         | item_id (int) in URL                            |                                                  |
| GET    | `/items`                 |                                                 | returns all rows                                 |
| POST   | `/signup`g               | `username` (string), `password` (argon2)        |                                                  |
| POST   | `/authenticate`          | `username` (string), `password` (argon2)        |                                                  |
| GET    | `/user/:user_id`         | user_id (int) in URL                            |                                                  |
| GET    | `/consumptions/:user_id` | user_id (int) in URL                            |                                                  |
| GET    | `/username               |                                                 | requires JWT token (obtain from `/authenticate`) |
| GET    | `/leaderboard/items`     |                                                 | returns 50 rows, sorted by # of consumptions     |
| GET    | `/leaderboard/users`     |                                                 | returns 50 rows, sorted by # of consumptions     |
