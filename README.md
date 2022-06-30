# Login Backend

## Development

### Getting started

Make sure you have the following installed:

- Docker / Docker desktop: https://www.docker.com/products/docker-desktop
- Go: https://go.dev
- GoLand: https://www.jetbrains.com/go/

> When you open the project with GoLand, it should prompt you to **install the required plugins**.

This repository contains **run configurations** which are loaded automatically when you open
the project in GoLand. To run the project, select the `go build` configuration in the top right
toolbar.

The `go build` configuration performs the following tasks for you:

- starts a postgres database with docker
- generates the Swagger code for the server
- generates the Prisma database client
- compiles & runs the project

### Migrating the database
When changes to the database are made, you can create a migration by running
```sh
POSTGRES_URL=postgresql://postgres:password@localhost:5432/postgres go run github.com/prisma/prisma-client-go migrate dev --name $MIGRATION_NAME
```
inside the project root dir.\
To apply all pending migrations, simply run
```sh
POSTGRES_URL=postgresql://postgres:password@localhost:5432/postgres go run github.com/prisma/prisma-client-go migrate deploy
```
inside the project root dir.

### Environment Variables
| Variable                | Description                                                                                                                                                                              |
|-------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `PORT`                  | OPTIONAL: You can use this to change the port that the API listens to. Default: 4000                                                                                                     |
| `POSTGRES_URL`          | **REQUIRED**: Connection URL to the postgres database. As an example, see [Migrating the database](#migrating-the-database) above.                                                       |
| `JWT_KEY`               | **REQUIRED**: We use this key to cryptographically sign the JWT token. Other backends can use this token to authenticate the user.                                                       |
| `JWT_VALIDITY_DURATION` | **REQUIRED**: The timespan how log each JWT is valid. The user has to log in again, if the token expires. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”. Example: "24h" |

### Project structure

| File / Directory                     | Description                                                                               |
|--------------------------------------|-------------------------------------------------------------------------------------------|
| `main.go`                            | Entry point of the program                                                                |
| `swagger/`                           | Contains the API specification (Swagger)                                                  |
| `prisma/`                            | Database schema (Prisma)                                                                  |
| `src/handlers/`                      | Our HTTP handler implementations (Go)                                                     |
| `src/handlers/handlers.go`           | Adds all our handler functions to the server                                              |
| `src/gen/`                           | Generated files from Prisma and Swagger. Do not edit or commit generated files!           |
| `src/gen/restapi/configure_login.go` | Configuration for the HTTP server. The only file in `src/gen/` that can be safely edited. |
