# Login Backend

## Development

### Getting started

Make sure you have the following installed:

- Go: https://go.dev
- GoLand: https://www.jetbrains.com/go/

You need to generate a few files, before you can start coding or compile the project.

```bash
go generate .
```

You should now be able to compile and run the server.

```bash
go run .
```

### Project setup

| File / Directory | Description                           |
|------------------|---------------------------------------|
| main.go          | Entry point of the program            |
| swagger/         | Contains the API specification        |
| gen/             | Contains generated files from swagger |
