# Backend

### Endpoints

**Note:** requests to the backend are limited to 100 requests/minute.

### V2

| Method | Route                                        | Parameters                                       | Description                                                                             | Notes                          |
| ------ | -------------------------------------------- | ------------------------------------------------ | --------------------------------------------------------------------------------------- | ------------------------------ |
| POST   | `/api/v2/items`                              | `name` (string), `units` (float)                 | add a new item                                                                          | requires authentication cookie |
| GET    | `/api/v2/items`                              |                                                  | get all items                                                                           |                                |
| GET    | `/api/v2/items/:name`                        | `name` (string) in URL                           | get a specific item from name                                                           |                                |
| GET    | `/api/v2/items/:name/leaderboard`            | `name` (string) in URL                           | get the top users for an item                                                           |                                |
| GET    | `/api/v2/items/:name/consumptions`           | `name` (string) in URL                           | get the number of consumptions for an item                                              |                                |
| POST   | `/api/v2/reports`                            | `name` (string), `reason` (string)               | report an item for having bad data, options are 'name', 'units', 'duplicate' or 'other' | requires authentication cookie |
| POST   | `/api/v2/consumptions`                       | `item_id` (int), `price` (float)                 | add a consumption                                                                       | requires authentication cookie |
| GET    | `/api/v2/consumptions/:consumption_id`       | `consumption_id` (int) in URL                    | get a consumption from its id                                                           | requires authentication cookie |
| DELETE | `/api/v2/consumptions`                       | `consumption_id` (int)                           | delete a consumption                                                                    | requires authentication cookie |
| GET    | `/api/v2/consumptions/count`                 |                                                  | get the total number of consumptions                                                    |                                |
| POST   | `/api/v2/signup`                             | `username` (string), `password` (string)         | create an account                                                                       |                                |
| POST   | `/api/v2/authenticate`                       | `username` (string), `password` (string)         | login to an account                                                                     | returns `HttpOnly` cookie      |
| POST   | `/api/v2/logout`                             |                                                  | logout from an account                                                                  | clears authentication cookie   |
| PUT    | `/api/v2/change_password`                    | `old_password` (string), `new_password` (string) | change password whilst logged in                                                        | requires authentication cookie |
| GET    | `/api/v2/users/:username`                    | `username` (string) in URL                       | get user information from username                                                      |                                |
| GET    | `/api/v2/users/me`                           |                                                  | get user information from logged in user                                                | requires authentication cookie |
| GET    | `/api/v2/users/:username/consumptions/count` | `username` (string) in URL                       | get the consumption count for a user                                                    |                                |
| GET    | `/api/v2/users/:username/consumptions`       | `username` (string) in URL                       | get the consumptions for a user (50 rows)                                               |                                |
| GET    | `/api/v2/users/:username/items/count`        | `username` (string) in URL                       | get the number of items a user has added to the database                                |
| GET    | `/api/v2/leaderboards/items`                 |                                                  | get the top 50 items ordered by consumptions                                            |                                |
| GET    | `/api/v2/leaderboards/users`                 |                                                  | get the top 10 users ordered by consumption count                                       |                                |
| GET    | `/api/v2/leaderboards/users/units`           |                                                  | get the top 10 users ordered by consumption units                                       |                                |
| GET    | `/api/v2/leaderboards/feed`                  |                                                  | get the most recent 10 consumptions                                                     |                                |

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
