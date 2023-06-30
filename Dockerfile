# Build Stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./createTestEnv main.go

# Run Stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/createTestEnv ./
RUN mkdir -p /app/data

CMD [ "/app/createTestEnv" ] 


