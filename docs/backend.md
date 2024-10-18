# Backend
* [Go](https://go.dev).
    * [Gin for HTTP API](https://pkg.go.dev/github.com/gin-gonic/gin).
    * [https://pkg.go.dev/github.com/lib/pq](pq for DB connection).

## Outline
### Users
#### User info
GIVE: user ID
GET:  user excl priv info)

### Items
#### Item info
GIVE: item ID
GET:  item

#### Items list
GIVE: <nothing>
GET:  item[]

* As items list grows, this function may become unsuitable due to network & client performance.
    * At this point either limit list (return n items at a time), or implement search func.

### Uses
#### Add
GIVE: user.id, timestamp
GET:  <success>

#### Get
GIVE: user.id, timeperiod
GET:  use[]

#### Getall
GIVE: user.id
GET:  use[]

* Again like [items list](####Items list) this might need to be optimised as the uses table grows.
