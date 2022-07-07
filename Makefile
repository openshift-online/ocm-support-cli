OUTPUT_DIR :=_output

# Constants
GOPATH := $(shell go env GOPATH)

build: clean
	go build -o ocm-support ./cmd/ocm-support || exit 1

install: clean
	go build -o $(GOPATH)/bin/ocm-support ./cmd/ocm-support || exit 1

clean:
	rm -f ocm-support


