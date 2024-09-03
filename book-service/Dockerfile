FROM golang:1.22.4-alpine3.20 as builder

WORKDIR /app
COPY . .

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/book-service
RUN go mod download && go mod tidy

RUN CGO_ENABLED=0 go build -o /usr/bin/book-service ./cmd

FROM alpine

COPY --from=builder /usr/bin/book-service /usr/bin/book-service
COPY --from=builder /app/book-service/migrations /app/book-service/migrations
COPY --from=builder /app/book-service/cmd/config.yml /app/book-service/cmd/config.yml

EXPOSE 9004

CMD ["book-service"]
