FROM golang:1.22.4-alpine3.20 as builder

WORKDIR /app
COPY . .

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/user-service
RUN go mod download && go mod tidy

WORKDIR /app/api-gateway
RUN go mod download && go mod tidy

RUN CGO_ENABLED=0 go build -o /usr/bin/api-gateway ./cmd

#RUN go build -o main .
FROM alpine

COPY --from=builder /usr/bin/api-gateway /usr/bin/api-gateway

EXPOSE 9000

CMD ["api-gateway"]
