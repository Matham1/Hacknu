# comp-prog-kz

Setup the project (Go language and dependencies should be already installed):
```bash
docker build -t go-image .
docker compose up
```

Project Overview: 
1. golang-migrate is installed for migration control. 
2. Server works with SQLC tool to compile and run database queries.

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