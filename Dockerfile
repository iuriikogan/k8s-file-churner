# syntax=docker/dockerfile:1
FROM golang:1.21-alpine As Builder

WORKDIR /app

RUN --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETOS TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /bin/k8sFileChurner main.go

FROM alpine:3 As Final

COPY --from=Builder /bin/k8sFileChurner /bin/

ENTRYPOINT ["/bin/k8sFileChurner"]
