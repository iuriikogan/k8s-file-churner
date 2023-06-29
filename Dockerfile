FROM golang:1.19-alpine

# Set destination for COPY
WORKDIR /

# Download Go modules
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/createTestEnv main.go
RUN mkdir -p /data

# Run
CMD [ "sh", "-c", "app/createTestEnv" ]


