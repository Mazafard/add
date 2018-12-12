PACKAGES := add
.PHONY: test cover  examples

test:
	go get -v -u github.com/smartystreets/goconvey/convey ;

	for package in $(PACKAGES) ; do \
		go test -v github.com/Mazafard/$$package ; \
	done ; \


cover:
	echo "mode: set" > profile.cov ; \
	for package in $(PACKAGES) ; do \
		go test -v -a -coverprofile=tmp.cov github.com/Mazafard/$$package ; \
		cat tmp.cov | grep -v "mode: set" >> profile.cov ; \
	done ; \
	rm tmp.cov ; \

EXAMPLES := $(shell ls examples/*.go | sed -e 's/examples\///')

examples:
	for example in $(EXAMPLES) ; do \
	    cd examples \
		go build -o $$example examples/$$example ; \
	done ; \
