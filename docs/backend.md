# Backend
* [Go](https://go.dev).
    * [Gin for HTTP API](https://pkg.go.dev/github.com/gin-gonic/gin).
    * [PGX for DB connection](https://github.com/jackc/pgx).

## Outline
Here is the requested information in markdown tables:

### Endpoints
| Method | Route                      | Notes                                        | Implemented |
|--------|----------------------------|----------------------------------------------| ------------|
| POST   | `/submit/item`             | will require auth (JWT)                      | **No**      |
| POST   | `/submit/consumption`      | will require auth (JWT)                      | **No**      |
| GET    | `/item/:item_id`           |                                              | Yes         |
| GET    | `/items`                   | returns all rows                             | Yes         |
| POST   | `/signup`                  |                                              | Yes         |
| GET    | `/user/:user_id`           |                                              | Yes         |
| GET    | `/leaderboard/items`       | returns 50 rows, sorted by # of consumptions | Yes         |
| GET    | `/leaderboard/users`       | returns 50 rows, sorted by # of consumptions | Yes         |

