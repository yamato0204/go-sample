# syntax=docker/dockerfile:1

FROM golang:1.23.0-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/main.go

FROM golang:1.23.0-alpine AS final

COPY migrations /migrations
COPY --from=builder /bin/api /bin/

EXPOSE 8080

ENTRYPOINT [ "/bin/api" ]
