GO			= go
BINARY		= testaustime

.PHONY: all lint test

install: build
	mv $(BINARY) $(HOME)/.local/bin/testaustime

build: test lint
	$(GO) build -o $(BINARY)

test:
	go test ./...
lint:
	golangci-lint run ./...

clean:
	rm -f $(BINARY)

