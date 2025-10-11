# Backend

### Endpoints

**Note:** requests to the backend are limited to 100 requests/minute.

### V2

| Method | Route                                        | Parameters                                       | Notes                          |
| ------ | -------------------------------------------- | ------------------------------------------------ | ------------------------------ |
| POST   | `/api/v2/items`                              | `name` (string), `units` (float)                 | requires authentication cookie |
| GET    | `/api/v2/items`                              |                                                  |                                |
| GET    | `/api/v2/items/:name`                        | `name` (string) in URL                           |                                |
| GET    | `/api/v2/items/:name/leaderboard`            | `name` (string) in URL                           |                                |
| GET    | `/api/v2/items/:name/consumptions`           | `name` (string) in URL                           |                                |
| POST   | `/api/v2/reports`                            | `name` (string), `reason` (string)               | requires authentication cookie |
| POST   | `/api/v2/consumptions`                       | `item_id` (int), `price` (float)                 | requires authentication cookie |
| GET    | `/api/v2/consumptions/:consumption_id`       | `consumption_id` (int) in URL                    | requires authentication cookie |
| DELETE | `/api/v2/consumptions`                       | `consumption_id` (int)                           | requires authentication cookie |
| GET    | `/api/v2/consumptions/count`                 |                                                  |                                |
| POST   | `/api/v2/signup`                             | `username` (string), `password` (string)         |                                |
| POST   | `/api/v2/authenticate`                       | `username` (string), `password` (string)         | returns `HttpOnly` cookie      |
| POST   | `/api/v2/logout`                             |                                                  | clears authentication cookie   |
| PUT    | `/api/v2/change_password`                    | `old_password` (string), `new_password` (string) | requires authentication cookie |
| GET    | `/api/v2/users/:username`                    | `username` (string) in URL                       |                                |
| GET    | `/api/v2/users/me`                           |                                                  | requires authentication cookie |
| GET    | `/api/v2/users/:username/consumptions/count` | `username` (string) in URL                       |                                |
| GET    | `/api/v2/users/:username/consumptions`       | `username` (string) in URL                       |                                |
| GET    | `/api/v2/users/:username/items/count`        | `username` (string) in URL                       |                                |
| GET    | `/api/v2/leaderboards/items`                 |                                                  |                                |
| GET    | `/api/v2/leaderboards/users`                 |                                                  |                                |
| GET    | `/api/v2/leaderboards/users/units`           |                                                  |                                |
| GET    | `/api/v2/leaderboards/feed`                  |                                                  |                                |


### V1

This will be deprecated at some point so don't use- there's nothing it does that V2 can't.

| Method | Route                                 | Parameters                                       | Notes                                                  |
| ------ | ------------------------------------- | ------------------------------------------------ | ------------------------------------------------------ |
| POST   | `/api/v1/submit/item`                 | `name` (string), `units` (int)                   | requires authentication cookie                         |
| POST   | `/api/v1/submit/consumption`          | `item_id` (int), `price` (float)                 | requires authentication cookie                         |
| POST   | `/api/v1/remove/consumption`          | `item_id` (int)                                  | requires authentication cookie                         |
| GET    | `/api/v1/items`                       |                                                  |                                                        |
| GET    | `/api/v1/items/:name`                 |                                                  |                                                        |
| GET    | `/api/v1/items/:name/leaderboard`     |                                                  |                                                        |
| GET    | `/api/v1/items/:name/consumptions`    |                                                  |                                                        |
| GET    | `/api/v1/items/:name/report`          |                                                  |                                                        |
| GET    | `/api/v1/consumption/:consumption_id` |                                                  | requires authentication cookie                         |
| POST   | `/api/v1/remove/consumption`          | `consumption_id` (int)                           | requires authentication cookie                         |
| POST   | `/api/v1/signup`                      | `username` (string), `password` (string)         |                                                        |
| POST   | `/api/v1/authenticate`                | `username` (string), `password` (string)         | returns `HttpOnly` cookie                              |
| POST   | `/api/v1/logout`                      |                                                  | clears authentication cookie                           |
| POST   | `/api/v1/change_password`             | `old_password` (string), `new_password` (string) | requires authentication cookie                         |
| GET    | `/api/v1/user/:username`              |                                                  |                                                        |
| GET    | `/api/v1/user/me`                     |                                                  | requires authentication cookie, returns logged in user |
| GET    | `/api/v1/consumption_count`           |                                                  |                                                        |
| GET    | `/api/v1/consumption_count/:username` |                                                  |                                                        |
| GET    | `/api/v1/consumptions/:username`      |                                                  |                                                        |
| GET    | `/api/v1/leaderboard/items`           |                                                  | returns 50 rows, sorted by # of consumptions           |
| GET    | `/api/v1/leaderboard/users`           |                                                  | returns 50 rows, sorted by # of consumptions           |
| GET    | `/api/v1/leaderboard/users-by-units`  |                                                  | returns 50 rows, sorted by # of units                  |
| GET    | `/api/v1/leaderboard/feed`            |                                                  | returns 10 rows of consumptions, sorted by timestamp   |
