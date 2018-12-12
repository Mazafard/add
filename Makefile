PACKAGES := add
.PHONY: test cover  examples

test:
	go get -v -u github.com/smartystreets/goconvey/convey ;

	for package in $(PACKAGES) ; do \
		go test -v github.com/Mazafard/$$package ; \
	done ; \


EXAMPLES := $(shell ls examples/*.go | sed -e 's/examples\///')

examples:
	for example in $(EXAMPLES) ; do \
	    cd examples \
		go build -o $$example examples/$$example ; \
	done ; \
