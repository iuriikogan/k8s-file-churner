#!/bin/bash
FROM golang:1.19-bookworm As Builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/k8sFileChurner main.go

FROM redhat/ubi9-micro
COPY --from=Builder /bin/k8sFileChurner /bin/k8sFileChurner
ENTRYPOINT ["/bin/k8sFileChurner"]
