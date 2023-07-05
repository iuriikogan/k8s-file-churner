#!/bin/bash
FROM golang:1.19-alpine
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/k8sFileChurner main.go
CMD ["/app/k8sFileChurner"]
