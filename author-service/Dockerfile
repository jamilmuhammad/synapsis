FROM golang:1.22.4-alpine3.20 as builder

WORKDIR /app
COPY . .

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/author-service
RUN go mod download && go mod tidy

RUN CGO_ENABLED=0 go build -o /usr/bin/author-service ./cmd

FROM alpine

COPY --from=builder /usr/bin/author-service /usr/bin/author-service
COPY --from=builder /app/author-service/migrations /app/author-service/migrations
COPY --from=builder /app/author-service/cmd/config.yml /app/author-service/cmd/config.yml

EXPOSE 9003

CMD ["author-service"]
