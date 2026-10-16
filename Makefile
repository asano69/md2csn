.PHONY: test convert

test:
	go test ./...

convert:
	go run . testdata/Universe.md
