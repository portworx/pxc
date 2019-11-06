CLINAME := pxc
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
VER := $(shell git describe --tags)
#VER := 0.0.0-$(SHA)
ARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)
DIR=.

ifdef APP_SUFFIX
  VERSION = $(VER)-$(subst /,-,$(APP_SUFFIX))
else
ifeq (master,$(BRANCH))
  VERSION = $(VER)
else
  VERSION = $(VER)-$(BRANCH)
endif
endif
LDFLAGS :=-ldflags "-X github.com/portworx/pxc/cmd.PxVersion=$(VERSION)"

ifneq (windows,$(GOOS))
PKG_NAME = $(CLINAME)
else
PKG_NAME = $(CLINAME).exe
endif

PACKAGE := $(CLINAME)-$(VERSION).$(GOOS).$(ARCH).zip

all: pxc

install:
	go install

imports:
	goimports -w ./cmd
	goimports -w ./handler
	goimports -w ./pkg
	goimports -w *.go

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

pxc:
	go build $(LDFLAGS)

release: darwin_amd64_dist \
	windows_amd64_dist \
	linux_amd64_dist

darwin_amd64_dist:
	GOOS=darwin GOARCH=amd64 $(MAKE) dist

windows_amd64_dist:
	GOOS=windows GOARCH=amd64 $(MAKE) dist

linux_amd64_dist:
	GOOS=linux GOARCH=amd64 $(MAKE) dist

dist: $(PACKAGE)

test:
	./hack/test.sh

verify: all test
	go fmt $(go list ./... | grep -v vendor) | wc -l | grep 0
	go vet $(go list ./... | grep -v vendor)

$(PACKAGE): all
	@echo Packaging client Binaries...
	@mkdir -p dist
	@zip dist/$@ $(PKG_NAME)
	@rm -f $(PKG_NAME)

clean:
	rm -f $(PKG_NAME)
	rm -rf dist

.PHONY: dist all clean darwin_amd64_dist windows_amd64_dist linux_amd64_dist \
	install release pxc test

