FROM golang:1.21-alpine

WORKDIR /server

COPY . .

RUN go build -o app cmd/main.go

EXPOSE 8080

CMD ["./app"]
