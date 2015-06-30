GO ?= godep go
COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
endif

all: build test cover
godep:
	go get github.com/tools/godep
godep-save:
	godep save ./...
build:
	if [ ! -d bin ]; then mkdir bin; fi
	$(GO) build -v -o bin/go-remote-config
fmt:
	$(GO) fmt ./...
test:
	if [ ! -d coverage ]; then mkdir coverage; fi
	$(GO) test -v ./ -race -cover -coverprofile=$(COVERAGEDIR)/remoteconfig.coverprofile
cover:
	$(GO) tool cover -html=$(COVERAGEDIR)/remoteconfig.coverprofile -o $(COVERAGEDIR)/remoteconfig.html
tc: test cover
coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)
clean:
	$(GO) clean
	rm -f bin/go-remote-config
	rm -rf coverage/
