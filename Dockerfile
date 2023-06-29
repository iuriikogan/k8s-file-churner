FROM golang:1.19-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY *.go ./

# Build
<<<<<<< HEAD
RUN CGO_ENABLED=0 GOOS=linux go build -o ./createTestEnv main.go
RUN chmod +x /app/createTestEnv
RUN chown -R 1001:1001 /app/createTestEnv
=======
RUN CGO_ENABLED=0 GOOS=linux go build -o createTestEnv main.go
>>>>>>> 541b92d (merge)
RUN mkdir -p /app/data
# Run
ENTRYPOINT ["/app/createTestEnv"]
