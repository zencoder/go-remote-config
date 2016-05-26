GO15VENDOREXPERIMENT := 1

COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
endif

all: build test cover
install-deps:
	glide install	
build:
	if [ ! -d bin ]; then mkdir bin; fi
	go build -v -o bin/go-remote-config
fmt:
	go fmt ./...

test: export AWS_ACCESS_KEY_ID := 1
test: export AWS_SECRET_ACCESS_KEY := 1
test:
	if [ ! -d coverage ]; then mkdir coverage; fi
	go test -v ./ -race -cover -coverprofile=$(COVERAGEDIR)/remoteconfig.coverprofile
cover:
	go tool cover -html=$(COVERAGEDIR)/remoteconfig.coverprofile -o $(COVERAGEDIR)/remoteconfig.html
tc: test cover
coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)
clean:
	go clean
	rm -f bin/go-remote-config
	rm -rf coverage/
