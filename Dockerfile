FROM golang:1.19-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o createTestEnv main.go
RUN mkdir -p /app/data
# Run
ENTRYPOINT ["/app/createTestEnv"]
