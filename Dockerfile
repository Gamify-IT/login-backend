FROM golang:1.18.1-bullseye AS build

# Setup working directory
WORKDIR /app

# Download dependencies into cache
COPY go.mod go.sum ./

RUN go mod download

# Download prisma binary
RUN go run github.com/prisma/prisma-client-go prefetch

# Generate libraries
COPY . .
RUN go generate

# Build
RUN go build -ldflags="-s -w" -o server .

# PORT for the REST API
ENV PORT=4000 HOST=0.0.0.0
EXPOSE $PORT

HEALTHCHECK --interval=5s --timeout=5s --start-period=10s --retries=2 CMD curl http://localhost:4000/health

# 1. Execute migrations
# 2. Run server
CMD go run github.com/prisma/prisma-client-go migrate deploy && ./server --port $PORT --host $HOST
