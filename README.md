# Informal Project
- Name tbd

## Team
1. M: backend, sysadmin
2. A: dba, data
3. E: full-stack(?), possibly also dba & data
4. C: requirements gathering and user testing

## Outline
* Data sourced from dbas and users.
    * This means the data will be sparse at first, but item uses will conform to a bell-curve. Once most of the commonly used items are added, adding new items will be rare.
    * L has good knowledge of items, could create a small tool for adding lots of new items.
* User creates account, adds items and records uses of items.

## Workflow
* Monorepo for backend (api), frontend(s) and db scripts (and laster on data analysis stuff).
* CI for baclend (gh actions).
* Use issues and PRs to keep track of work.
    * Review PRs before merge: increase understanding of project.
* Use [TDD](https://en.wikipedia.org/wiki/Test-driven\_development).
## Plan
**No time restraint on stages**
1. DB & backend implementation of users and items.
2. DB & backend implementation of uses of items.
3. Frontend implementations (web, bot).
4. Data analysis on uses.

Here is a prettier version of your markdown content:

## Stack

| **Service**   | **Software**   |
|---------------|----------------|
| **DB**        | PostgreSQL     |
| **Backend**   | Go             |
| **Frontend**  | HTML / Bot     |

---

## DB

- **[PostgreSQL](https://postgresql.org)**.
- **Tables**:
  1. **Items**
  2. **Users**
  3. **Consumptions**

---

## Backend

- **[Go](https://go.dev)**.
  - **[Gin for HTTP API](https://pkg.go.dev/github.com/gin-gonic/gin)**.
  - **[PGX for DB connection](https://github.com/jackc/pgx)**.
    - **[Mock tests (for TDD)](https://github.com/jackc/pgmock)**.
- More details in the **[Backend Specification](./docs/backend.md)**.

---

## Frontend

- **Web Frontend**:
  - Basic **HTML + JS** should suffice.
  - Ensure it performs well on **mobile devices**.
- **Discord bot** built with **Python 3**.
- More details in the **[Frontend Specification](./docs/frontend.md)**.
