FROM golang:1.21-alpine

WORKDIR /server

RUN go install github.com/cosmtrek/air@latest

COPY . .

# RUN go build -o app cmd/main.go
RUN go mod download

EXPOSE 8080

# CMD ["./app"]
CMD ["air", "-c", ".air.toml"]
