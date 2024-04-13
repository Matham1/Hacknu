FROM golang:1.21.6-alpine3.18

WORKDIR /app

COPY . . 

RUN go get -d -v ./...

RUN go build -o api -v ./cmd/apiserver/main.go

EXPOSE 8080

CMD ["./api"]
