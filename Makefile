.PHONY: all
all: test lint demand

.PHONY: clean
clean:
	rm -fR bin
	
.PHONY: test
test:
	go test -v ./...

.PHONY: demand
demand:
	mkdir -p bin
	go build -o bin .

.PHONY: lint
lint: bin/tools/golangci-lint
	bin/tools/golangci-lint run

bin/tools/golangci-lint:
	mkdir -p bin/tools
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "bin/tools" v1.52.2
