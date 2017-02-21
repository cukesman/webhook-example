.PHONY: default
default: build;

get-deps:
	echo "No deps"

build:
	go build -o ./build/webhook .

run:
	./build/webhook

clean:
	rm -rf ./build
