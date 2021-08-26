build:
	go build -o=bin/trab1 src/*.go
run: build
	./bin/trab1