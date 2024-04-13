# comp-prog-kz

Setup

Setup the rabbitmq docker container by using the Dockerfile with the following command
```bash
docker build -t rabbitmq .
docker run -d --name rabbitmq-server -p 5672:5672 -p 15672:15672 rabbitmq
```

To run the code, write the following command:
```bash
$ go run -v ./...
```

To build the code, write the following command:
```bash
$ go build -v ./...
```

To test the code, write the following command:
```bash
$ go test -v ./...
```