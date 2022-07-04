# Login Backend

## IMPORTANT

This project depends on a lot of generated code.
When you first clone the project, you need to either follow the [Getting started](#getting-started) guide, or run `go generate`.
Every time you change something in `swagger/swagger.yaml` or `prisma/schema.prisma`, you **must** regenerate the files for your changes to take effect.
Do this by again using the `go generate` command or the `go build` run configuration.

## Development

Here is how you get started when developing on the login backend.

### Getting started

Make sure you have the following installed:

- Docker / Docker desktop: https://www.docker.com/products/docker-desktop
- Go: https://go.dev
- GoLand: https://www.jetbrains.com/go/

> When you open the project with GoLand, it should prompt you to **install the required plugins**.

The project contains two run configurations which GoLand will automatically load.
You can find them in the top right toolbar.

First, you need to start a database. Select "run database" in the dropdown menu (left of the green play button).
Then click the play button to start the database.

Initially, the database is empty. You need to create the database tables with the following command:
```sh
POSTGRES_URL=postgresql://postgres:password@localhost:5432/postgres go run github.com/prisma/prisma-client-go migrate deploy
```

Finally, to run the project, select the `go build` configuration in the dropdown menu (left of the green play button).
Click the play button to compile and run the project.

The server prints the URL where it is running to the console.
Visit http://localhost:4000/docs to see the API documentation.

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
