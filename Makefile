.PHONY: 
	update
	generate
	build

update:
	go mod tidy

generate:
	go generate ./...

build:
	go build -o accounts