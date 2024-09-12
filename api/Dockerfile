# syntax=docker/dockerfile:1

# ARG GO_VERSION=1.23.0

FROM golang:1.23.0-alpine AS builder
WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/main.go

FROM golang:1.23.0-alpine as final

COPY migrations /migrations
COPY --from=builder /bin/api /bin/

EXPOSE 8080

ENTRYPOINT [ "/bin/api" ]
