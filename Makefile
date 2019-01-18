.PHONY: 
	update
	generate
	build
	clean

update:
	go mod tidy

generate:
	go generate ./...

build:
	go build -o accounts

clean:
	rm -f accounts

run:
	go run main.go