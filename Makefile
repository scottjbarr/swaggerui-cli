.PHONY: dist

dist: prepare build

prepare:
	mkdir -p tmp dist

build:
	go build -o dist/swaggerui-cli cmd/swaggerui-cli/main.go

clean:
	rm -rf dist
