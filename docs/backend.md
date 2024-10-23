# Backend
* [Go](https://go.dev).
    * [Gin for HTTP API](https://pkg.go.dev/github.com/gin-gonic/gin).
    * [PGX for DB connection](https://github.com/jackc/pgx).

## Outline
Here is the requested information in markdown tables:

### Users
| **Action**   | **Input (GIVE)**   | **Output (GET)**           | **Relative URL**    |
|--------------|--------------------|----------------------------|---------------------|
| User info    | User ID             | User excl priv info         | `/users/{user_id}`     |

### Items
| **Action**    | **Input (GIVE)**   | **Output (GET)**   | **Relative URL**     |
|---------------|--------------------|--------------------|----------------------|
| Item info     | Item ID             | Item               | `/item/{item_id}`      |
| Items list    | <nothing>           | Item[]             | `/items`                |

> **Note:** As the items list grows, this function may become unsuitable due to network & client performance. At this point, either limit the list (return *n* items at a time), or implement a search function.

### Consumptions
| **Action** | **Input (GIVE)**         | **Output (GET)**   | **Relative URL**      |
|------------|--------------------------|--------------------|-----------------------|
| Add        | User ID, Item ID          | <success>          | `/consumptions/add`      |
| Get        | User ID, Time period(?)   | Use[]              | `/consumptions/{user_id}?period={time_period}` |
| Get all(?) | User ID                   | Use[]              | `/consumptions/{user_id}` |

> **Note:** Like the Items list, this function might need to be optimized as the Uses table grows.

