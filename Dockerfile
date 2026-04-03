FROM golang:1.24-alpine AS builder

ENV GOTOOLCHAIN=auto

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .
RUN swag init -g cmd/app/main.go -o docs
RUN CGO_ENABLED=0 go build -o /bin/app ./cmd/app

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/app /bin/app
COPY --from=builder /go/bin/migrate /bin/migrate
COPY migrations /migrations
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
