FROM golang:1.19-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./createTestEnv main.go
RUN chmod +x /app/createTestEnv
RUN chown -R 1001:1001 /app/createTestEnv
RUN mkdir -p /app/data
# Run
ENTRYPOINT ["/app/createTestEnv"]
