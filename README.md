# Swagger UI CLI

Run a Swagger UI server.

Uses the github.com/flowchartsman/swaggerui package.

## Usage

```
Usage of ./dist/swaggerui-cli
  -bind string
    	Address to bind to. Default localhost:8080 (default "localhost:8080")
  -filename string
    	Swagger file to load
  -prefix string
    	Path prefix . Defaults to "/" (default "/")
```

## Example

```
$ go run cmd/swaggerui-cli/main.go -filename tmp/pet-store.json -prefix "/" -bind localhost:8080
```

## Build

Build the binary for your local machine.

```
$ make
```

## License

The MIT License (MIT)

Copyright (c) 2024 Scott Barr

See [LICENSE](LICENSE)
