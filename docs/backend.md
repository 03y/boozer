# Backend
* [Go](https://go.dev).
    * [Gin for HTTP API](https://pkg.go.dev/github.com/gin-gonic/gin).
    * [https://pkg.go.dev/github.com/lib/pq](pq for DB connection).

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

### Uses
| **Action** | **Input (GIVE)**         | **Output (GET)**   |
|------------|--------------------------|--------------------|
| Add        | User ID, Timestamp        | <success>          |
| Get        | User ID, Time period      | Use[]              |
| Getall     | User ID                   | Use[]              |

> **Note:** Like the Items list, this function might need to be optimized as the Uses table grows.
