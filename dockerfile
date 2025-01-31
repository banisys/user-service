FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache git curl && \
    go install github.com/air-verse/air@latest && \
    go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

ENV GOFLAGS="-buildvcs=false"
ENV CGO_ENABLED=1

RUN apk add --no-cache gcc musl-dev

CMD ["air"]