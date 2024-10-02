format:
	gofmt -s -w .

build:
	$(MAKE) format
	go build -o bin/ggit main.go
	chmod 770 bin/ggit

run:
	go run main.go