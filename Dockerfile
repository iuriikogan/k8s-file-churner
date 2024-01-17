# syntax=docker/dockerfile:1
FROM golang:1.19-bookworm As Builder

WORKDIR /app

COPY go.* ./

RUN go mod download -x

COPY . ./

RUN --mount=type=cache,target=/go/pkg/mod/,--mount=type=bind,target=. \
CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/k8sFileChurner main.go

FROM scratch as Final

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    churner
USER churner

COPY --from=Builder /bin/k8sFileChurner /bin/

ENTRYPOINT ["/bin/k8sFileChurner"]
