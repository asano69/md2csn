.PHONY: test convert build

test:
	go test ./...

convert:
	go run . testdata/Universe.md

build:
	go build -o md2csn .
