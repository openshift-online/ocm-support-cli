OUTPUT_DIR :=_output

# Constants
GOPATH := $(shell go env GOPATH)

build: clean
	go build -o ocm-support ./cmd/ocm-support || exit 1

install: clean
	go build -o $(GOPATH)/bin/ocm-support ./cmd/ocm-support || exit 1

clean:
	rm -f ocm-support

cmds:
	for cmd in $$(ls cmd); do \
		go build "./cmd/$${cmd}" || exit 1; \
	done

ensureOCM:
	bash ensure_ocm_cli.sh

test:
	for cmd in $$(ls cmd); do \
		go build "./cmd/$${cmd}" || exit 1; \
	done
	bash ensure_ocm_cli.sh
	ginkgo run -r