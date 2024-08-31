FROM golang:1.22.4-alpine3.20 as builder

WORKDIR /app
COPY . .

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/category-service
RUN go mod download && go mod tidy

RUN CGO_ENABLED=0 go build -o /usr/bin/category-service ./cmd

FROM alpine

COPY --from=builder /usr/bin/category-service /usr/bin/category-service
COPY --from=builder /app/category-service/migrations /app/category-service/migrations
COPY --from=builder /app/category-service/cmd/config.yml /app/category-service/cmd/config.yml

EXPOSE 9002

CMD ["category-service"]
