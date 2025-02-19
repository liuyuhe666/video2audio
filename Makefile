all: build test

build:
	go build -o video2audio

test:
	go test ./

clean:
	rm -f video2audio