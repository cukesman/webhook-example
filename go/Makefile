.PHONY: default
default: build

all: clean get-deps build test

get-deps:
	go get "github.com/stretchr/testify/assert"

build:
	go build -o ./build/webhook -ldflags "-linkmode external -extldflags -static" .

run:
	./build/webhook

test:
	go test -v .

clean:
	rm -rf ./build
