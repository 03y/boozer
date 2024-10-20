# Backend
* [Go](https://go.dev).
    * [Gin for HTTP API](https://pkg.go.dev/github.com/gin-gonic/gin).
    * [PGX for DB connection](https://github.com/jackc/pgx)**.

## Outline
Here is the requested information in markdown tables:

### Users
| **Action**   | **Input (GIVE)**   | **Output (GET)**           |
|--------------|--------------------|----------------------------|
| User info    | User ID             | User excl priv info         |

### Items
| **Action**    | **Input (GIVE)**   | **Output (GET)**   |
|---------------|--------------------|--------------------|
| Item info     | Item ID             | Item               |
| Items list    | <nothing>           | Item[]             |

> **Note:** As the items list grows, this function may become unsuitable due to network & client performance. At this point, either limit the list (return *n* items at a time), or implement a search function.

### Consumptions
| **Action** | **Input (GIVE)**         | **Output (GET)**   |
|------------|--------------------------|--------------------|
| Add        | User ID, Item ID          | <success>          |
| Get        | User ID, Time period(?)   | Use[]              |
| Get all(?) | User ID                   | Use[]              |

> **Note:** Like the Items list, this function might need to be optimized as the Uses table grows.
