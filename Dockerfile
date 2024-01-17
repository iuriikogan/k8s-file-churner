# syntax=docker/dockerfile:1
FROM golang:1.19-bookworm As Builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN --mount=type=cache,target=/go/pkg/mod/ \
CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/k8sFileChurner main.go

FROM redhat/ubi9-micro as Runner

COPY --from=Builder /bin/k8sFileChurner /bin/

ENTRYPOINT ["/bin/k8sFileChurner"]
