# 🍺 Boozer
Boozer is a web app which stores users, beers and consumptions!

## Outline
- Data sourced from users.
    - This means the data will be sparse at first, but item uses will conform to a bell-curve. Once most of the commonly used items are added, adding new items will be rare.
- User creates account, adds items and records consumptions.

## Team
1. [03y](https://github.com/03y): full-stack, sysadmin
2. [Adam-W2](https://github.com/Adam-W2): frontend(?), data analysis
3. [EuanRobertson1](https://github.com/EuanRobertson1): full-stack(?)
4. [Choob303](https://github.com/Choob303): requirements gathering and user testing

## Workflow
- Monorepo for backend (api), frontend and db scripts (and laster on data analysis stuff).
- CI for backend (gh actions).
- Use issues and PRs to keep track of work.
- Use behaviour-driven development (BDD) style testing

## Plan
**No time restraint on stages**
- [x] DB & backend implementation of users and items.
- [x] DB & backend implementation of uses of items.
- [x] Frontend implementation.
- [ ] Data analysis.

## Stack
| Service  | Software   |
|----------|------------|
| DB       | PostgreSQL |
| Backend  | Go         |
| Frontend | HTML/JS    |

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
  - **BDD tests with [hurl](https://hurl.dev)**.
- More details in the **[Backend Specification](./docs/backend.md)**.

---

## Frontend
- **Web Frontend**:
  - Basic **HTML + JS** (possibly look into use of JS framework).
  - Ensure it performs well on **mobile devices**.
- More details in the **[Frontend Specification](./docs/frontend.md)**.

## Running
The easiest way to run the entire application stack is with Docker, see the [backend documentation](./backend/README.md).
