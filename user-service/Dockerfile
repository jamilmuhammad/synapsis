FROM golang:1.22.4-alpine3.20 as builder

WORKDIR /app
COPY go.mod go.sum ./
COPY ./ /app/user-service/

#COPY go.mod go.sum ./
WORKDIR /app/user-service
RUN go mod download
RUN CGO_ENABLED=0 go build -o /usr/bin/user-service .

#COPY . .

#RUN go build -o main .
FROM alpine

EXPOSE 50051

COPY --from=builder /usr/bin/user-service /usr/bin/user-service

CMD ["user-service"]
