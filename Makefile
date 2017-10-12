COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
endif

ifdef VERBOSE
V = -v
else
.SILENT:
endif

install-deps:
	glide install
build:
	mkdir -p bin
	go build $(V) -o bin/go-remote-config
fmt:
	go fmt ./...

test: export AWS_ACCESS_KEY_ID := 1
test: export AWS_SECRET_ACCESS_KEY := 1
test:
	mkdir -p coverage
	go test $(V) ./ -race -cover -coverprofile=$(COVERAGEDIR)/remoteconfig.coverprofile
cover:
	go tool cover -html=$(COVERAGEDIR)/remoteconfig.coverprofile -o $(COVERAGEDIR)/remoteconfig.html
coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)
clean:
	go clean
	rm -f bin/go-remote-config
	rm -rf coverage/
