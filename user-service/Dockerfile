FROM golang:1.22.4-alpine3.20 as builder

WORKDIR /app
COPY . .

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/user-service
RUN go mod download && go mod tidy

RUN CGO_ENABLED=0 go build -o /usr/bin/user-service ./cmd

#RUN go build -o main .
FROM alpine

COPY --from=builder /usr/bin/user-service /usr/bin/user-service
COPY --from=builder /app/user-service/migrations /app/user-service/migrations
COPY --from=builder /app/user-service/cmd/config.yml /app/user-service/cmd/config.yml

EXPOSE 9001

CMD ["user-service"]
