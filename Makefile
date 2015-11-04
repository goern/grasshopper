# This is just a wrapper for the old school guys

OUT_DIR=_output
OUT_PKG_DIR=Godeps/_workspace/pkg

all build: main.go
	godep restore
	CGO_ENABLED=0 go build --ldflags '-extldflags "-static"'

test:
	go test -v github.com/goern/grasshopper/nulecule

image: build test
	strip grasshopper
	docker build --rm --tag goern/grasshopper:0.0.5 -f Dockerfile .

clean:
	rm -rf grasshopper

clean-image:
	docker rmi goern/grasshopper:0.0.5
