FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

# Final stage
FROM golang:1.23.3

WORKDIR /app

COPY --from=builder /go/bin/air /usr/local/bin/air
COPY . .

EXPOSE 8080

CMD ["air"]