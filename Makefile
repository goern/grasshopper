# This is just a wrapper for the old school guys

OUT_DIR = _output
OUT_PKG_DIR = Godeps/_workspace/pkg

all build: main.go
	CGO_ENABLED=0 go build --ldflags '-extldflags "-static"'

image: build
	strip grasshopper
	docker build --rm --tag goern/grasshopper:0.0.3 -f Dockerfile .

clean:
	rm -rf grasshopper

clean-image:
	docker rmi goern/grasshopper:0.0.3
