SHELL 		= /bin/bash
GO			= go
BINARY		= testaustime
identifier	= testaustime-cli

.PHONY: all lint test

install: build setup
beforecommit: build clean
	go mod tidy
	go fmt ./...

uninstall: clean
	rm -rf \
		$(HOME)/.config/$(identifier) \
		$(HOME)/.local/share/$(identifier) \
		$(HOME)/.local/bin/$(BINARY)

setup:
	mkdir -p \
		$(HOME)/.config/$(identifier) \
		$(HOME)/.local/share/$(identifier)
	cp ./example.toml ~/.config/$(identifier)/config.toml
	mv $(BINARY) $(HOME)/.local/bin

build: test 
	$(GO) build -o $(BINARY)

test:
	go test ./...

lint:
	golangci-lint run ./...

clean:
	rm -f $(BINARY)

