FROM golang:1.24-alpine

ENV GOTOOLCHAIN=auto

RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY air.entrypoint.sh /air.entrypoint.sh
RUN chmod +x /air.entrypoint.sh

WORKDIR /app

ENTRYPOINT ["/air.entrypoint.sh"]
