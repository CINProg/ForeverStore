build:
	@go build -o bin/fs

run: build
	@./bin/fs

test:
	@go tets ./... -v