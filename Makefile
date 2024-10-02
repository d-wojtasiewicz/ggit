clean:
	rm -rf ./bin || true

format:
	gofmt -s -w .

build:
	$(MAKE) clean || true
	$(MAKE) format
	go build -o bin/ggit main.go
	chmod 770 bin/ggit

run:
	go run main.go