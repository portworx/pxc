CLINAME := pxc
KUBECTL_PLUGIN := kubectl-$(CLINAME)
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
VER := $(shell git describe --tags)
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
PLUGIN_PKG_NAME = $(KUBECTL_PLUGIN)
else
PKG_NAME = $(CLINAME).exe
PLUGIN_PKG_NAME = $(KUBECTL_PLUGIN).exe
endif

ZIPPACKAGE := $(CLINAME)-$(VERSION).$(GOOS).$(ARCH).zip
TGZPACKAGE := $(CLINAME)-$(VERSION).$(GOOS).$(ARCH).tar.gz

all: pxc $(PLUGIN_PKG_NAME)

install: all
	cp $(PKG_NAME) $(GOPATH)/bin
	cp $(PKG_NAME) $(GOPATH)/bin/$(PLUGIN_PKG_NAME)

imports:
	goimports -w ./cmd
	goimports -w ./handler
	goimports -w ./pkg
	goimports -w *.go

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

pxc:
	go build -o $(PKG_NAME) $(LDFLAGS)

release: darwin_amd64_dist \
	windows_amd64_dist \
	linux_amd64_dist

darwin_amd64_dist:
	GOOS=darwin GOARCH=amd64 $(MAKE) dist

windows_amd64_dist:
	GOOS=windows GOARCH=amd64 $(MAKE) distzip

linux_amd64_dist:
	GOOS=linux GOARCH=amd64 $(MAKE) dist

distzip: $(ZIPPACKAGE)

dist: $(TGZPACKAGE)

# This also tests for any conflicts
docs: all
	./pxc gendocs --output-dir=docs/usage

test:
	./hack/test.sh

verify: all test
	go fmt $(go list ./... | grep -v vendor) | wc -l | grep 0
	go vet $(go list ./... | grep -v vendor)

$(PLUGIN_PKG_NAME): pxc
	cp pxc $(PLUGIN_PKG_NAME)

$(ZIPPACKAGE): all
	@echo Packaging pxc ...
	@mkdir -p dist
	@zip dist/$@ $(PLUGIN_PKG_NAME)
	@rm -f $(PKG_NAME)

$(TGZPACKAGE): all
	@echo Packaging Binaries...
	@mkdir -p tmp/$(PKG_NAME)
	@cp $(PLUGIN_PKG_NAME) tmp/$(PKG_NAME)/
	@mkdir -p $(DIR)/dist/
	tar -czf $(DIR)/dist/$@ -C tmp $(PKG_NAME)
	@rm -rf tmp

clean:
	rm -f $(PKG_NAME) $(PLUGIN_PKG_NAME)
	rm -rf dist

.PHONY: dist all clean darwin_amd64_dist windows_amd64_dist linux_amd64_dist \
	install release pxc test

