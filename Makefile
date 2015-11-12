# This is just a wrapper for the old school guys

OUT_DIR=_output
OUT_PKG_DIR=Godeps/_workspace/pkg
GRASSHOPPER_MIN_VERSION=$(shell date -u +%Y%m%d.%H%M%S)

.PHONY: all build
all build: main.go
	godep restore
	CGO_ENABLED=0 go build --ldflags '-extldflags "-static" -X github.com/goern/grasshopper/cmd.minversion=$(GRASSHOPPER_MIN_VERSION) -X github.com/goern/grasshopper/cmd.version=$(GRASSHOPPER_VERSION)'

.PHONY: test
test:
	go test -v github.com/goern/grasshopper/nulecule

.PHONY: image
image: build test
	strip grasshopper
	docker build --rm --tag goern/grasshopper:$(GRASSHOPPER_VERSION) -f Dockerfile .

.PHONY: doc
doc:
	asciidoc --backend=html5 README.asciidoc
	a2x -d manpage -f manpage grasshopper.8.asciidoc

.PHONY: clean
clean:
	rm -rf grasshopper

.PHONY: clean-image
clean-image:
	docker rmi goern/grasshopper:$(GRASSHOPPER_VERSION)
