# Card Deck API
Built with Go 1.17.3

## Running the API
1. Open the main project directory in a terminal
2. Run `go run ./main.go`
3. The project should start running at `localhost:13370` (or `127.0.0.1:13370`)
4. Swagger is available at `localhost:13370/swagger/index.html`

## Running tests
1. Open the main project directory in a terminal
2. Run `go test ./...`

## Troubleshooting
You might encounter an error if something else in your system uses the **13370** port.

You can run the API at different port, by specifying it in a `port` flag, e.g:
```
go run ./main.go -port 5077
```
