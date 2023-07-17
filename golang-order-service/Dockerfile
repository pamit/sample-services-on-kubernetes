FROM golang:1.20-alpine

RUN apk update && apk add --no-cache git

WORKDIR /order_service
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go mod tidy
RUN go build -o /golang-order-service

EXPOSE 8000

CMD [ "/golang-order-service" ]
