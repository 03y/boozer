## Building and Running

### Manual

To build and run the backend manually, you'll need a Go compiler. You'll also need to have the DB running (see [../db/README.md](../db/README.md))

1.  **Build the application:**
    `go build`

2.  **Set the `DATABASE_URL` environment variable:**
    `export DATABASE_URL='postgres://postgres:postgres@localhost:5432/boozer'`

3.  **Run the application:**
    `./boozer 0.0.0.0:6969`

### Docker Compose

The easiest way to get the backend running is to use Docker Compose. This will build the backend and a PostgreSQL database and run them in containers.

1.  **Build and start the containers:**
    `docker-compose up --build` (add `-d` to run in background)

2.  **View the logs:**
    `docker-compose logs`

3.  **Stop the containers:**
    `docker-compose down`
