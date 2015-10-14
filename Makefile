# This is just a wrapper for the old school guys

GOPATH = $(shell pwd)
OUT_DIR = _output
OUT_PKG_DIR = Godeps/_workspace/pkg

all build:
	./build.sh $(WHAT)
.PHONY: all build

image: build
	docker build --rm --tag goern/grasshopper:0.0.3 -f Dockerfile .

clean:
	rm -rf grasshopper
	docker rmi goern/grasshopper:0.0.3
