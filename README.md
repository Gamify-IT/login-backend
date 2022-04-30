# Login Backend

## Development

### Getting started

Make sure you have the following installed:

- Go: https://go.dev
- GoLand: https://www.jetbrains.com/go/

> When you open the project with GoLand, it should prompt you to install the required plugins.

You need to generate a few files, before you can start coding or compile the project.

```bash
go generate
```

You should now be able to compile and run the server.

```bash
go run .
```

### Project setup

| File / Directory | Description                              |
|------------------|------------------------------------------|
| main.go          | Entry point of the program               |
| swagger/         | Contains the API specification (Swagger) |
| prisma/          | Database schema (Prisma)                 |
| src/handlers/    | Our HTTP handler implementation (Go)     |
| src/gen/         | Generated files from Prisma and Swagger  |
