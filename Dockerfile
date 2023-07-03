# Build Stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./k8s-file-churner main.go

# Run Stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/k8s-file-churner ./
RUN mkdir -p /app/data

ENTRYPOINT [ "/app/k8s-file-churner" ] 


