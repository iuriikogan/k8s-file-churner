#!/bin/bash
FROM golang:1.19-alpine As Builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /k8sFileChurner main.go

FROM alpine
COPY --from=Builder /k8sFileChurner /bin/k8sFileChurner
ENTRYPOINT ["/bin/k8sFileChurner"]
