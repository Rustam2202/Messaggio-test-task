FROM golang:1.22.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .env

RUN go build -o /message-processor

EXPOSE 8080

CMD ["/message-processor"]
