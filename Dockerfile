FROM golang:1.21.5 as base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM base as dev

COPY .env ./

CMD ["go", "run", "cmd/main.go"]