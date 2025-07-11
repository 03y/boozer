# Backend

- [Go](https://go.dev).
  - [Gin for HTTP API](https://pkg.go.dev/github.com/gin-gonic/gin).
  - [PGX for DB connection](https://github.com/jackc/pgx).

### Endpoints

| Method | Route                    | Parameters                                      | Notes                                            |
| ------ | ------------------------ | ----------------------------------------------- | ------------------------------------------------ |
| POST   | `/submit/item`           | `name` (string), `type` (string), `abv` (float) | requires authentication cookie                   |
| POST   | `/submit/consumption`    | `item_id` (int), `amount` (int)                 | requires authentication cookie                   |
| GET    | `/item/:item_id`         | item_id (int) in URL                            |                                                  |
| GET    | `/items`                 |                                                 | returns all rows                                 |
| POST   | `/signup`                | `username` (string), `password` (string)        |                                                  |
| POST   | `/authenticate`          | `username` (string), `password` (string)        | returns `HttpOnly` cookie                        |
| POST   | `/logout`                |                                                 | clears authentication cookie                     |
| GET    | `/user/:user_id`         | user_id (int) in URL                            |                                                  |
| GET    | `/consumptions/:user_id` | user_id (int) in URL                            |                                                  |
| GET    | `/username`              |                                                 | requires authentication cookie                   |
| GET    | `/leaderboard/items`     |                                                 | returns 50 rows, sorted by # of consumptions     |
| GET    | `/leaderboard/users`     |                                                 | returns 50 rows, sorted by # of consumptions     |
