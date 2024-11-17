# Backend
* [Go](https://go.dev).
    * [Gin for HTTP API](https://pkg.go.dev/github.com/gin-gonic/gin).
    * [PGX for DB connection](https://github.com/jackc/pgx).

## Outline
Here is the requested information in markdown tables:

### Endpoints
| Method | Route                      | Notes                                            | Implemented |
|--------|----------------------------|--------------------------------------------------|-------------|
| POST   | `/submit/item`             | requires JWT token (obtain from `/authenticate`) | Yes         |
| POST   | `/submit/consumption`      | requires JWT token (obtain from `/authenticate`) | Yes         |
| GET    | `/item/:item_id`           |                                                  | Yes         |
| GET    | `/items`                   | returns all rows                                 | Yes         |
| POST   | `/signup`                  |                                                  | Yes         |
| POST   | `/authenticate`            |                                                  | Yes         |
| GET    | `/user/:user_id`           |                                                  | Yes         |
| GET    | `/leaderboard/items`       | returns 50 rows, sorted by # of consumptions     | Yes         |
| GET    | `/leaderboard/users`       | returns 50 rows, sorted by # of consumptions     | Yes         |

