# DB
* [Postgresql](https://postgresql.org)

## Outline
Three tables: users, items and uses.

### Users
| Field         | Datatype      | Key | Comments  |
| ------------- |:-------------:|:---:|:---------:|
| user\_id      | int(20)       | PK  | AUTO-INC  |
| username      | varchar(20)   |     |           |
| password      | varchar(256)  |     | argon2    |
| created       | int           |     | unix time |

### Items
| Field         | Datatype      | Key | Comments  |
| ------------- |:-------------:|:---:|:---------:|
| item\_id      | int(20)       | PK  | AUTO-INC  |
| name          | varchar(40)   |     |           |
| units         | float         |     |           |
| added         | int           |     | unix time |

### Consumptions
| Field         | Datatype      | Key | Comments  |
| ------------- |:-------------:|:---:|:---------:|
| consumption\_id       | int(20)       | PK  | AUTO-INC  |
| item\_id      | int(20)       | FK  |           |
| user\_id      | int(20)       | FK  |           |
| time          | int           |     | unix time |
