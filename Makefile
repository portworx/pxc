CLINAME := pxc
KUBECTL_PLUGIN := kubectl-$(CLINAME)
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
VER := $(shell git describe --tags)
ARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)
PXC_LDFLAGS =
PXC_GOBUILD_FLAGS =
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
LDFLAGS := -ldflags "-X github.com/portworx/pxc/cmd.PxVersion=$(VERSION) $(PXC_LDFLAGS)"

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
	go build $(PXC_GOBUILD_FLAGS) -o $(PKG_NAME) $(LDFLAGS)

docker-release: darwin_amd64_dist \
	windows_amd64_dist \
	linux_amd64_dist

release:
	docker run --privileged -ti \
		-v $(shell pwd):/go/src/github.com/portworx/pxc \
		-w /go/src/github.com/portworx/pxc \
		-e DEV_USER=$(shell id -u) \
		-e DEV_GROUP=$(shell id -g) \
		golang \
		hack/create-release.sh

darwin_amd64_dist:
	GOOS=darwin GOARCH=amd64 $(MAKE) PXC_LDFLAGS="-s -w" dist

windows_amd64_dist:
	GOOS=windows GOARCH=amd64 $(MAKE) PXC_LDFLAGS="-s -w" distzip

linux_amd64_dist:
	GOOS=linux GOARCH=amd64 $(MAKE) PXC_LDFLAGS="-s -w" dist

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
	$(MAKE) -C component/examples/golang verify

$(PLUGIN_PKG_NAME): pxc
	cp $(PKG_NAME) $(PLUGIN_PKG_NAME)

$(ZIPPACKAGE): all
	@echo Packaging pxc ...
	@mkdir -p tmp/$(PKG_NAME)
	@cp $(PLUGIN_PKG_NAME) tmp/$(PKG_NAME)/
	@cp extras/docs/* tmp/$(PKG_NAME)/
	@mkdir -p $(DIR)/dist
	@( cd tmp/$(PKG_NAME) ; zip ../../dist/$@ * )
	@rm -f $(PKG_NAME) $(PLUGIN_PKG_NAME)
	@rm -rf tmp

$(TGZPACKAGE): all
	@echo Packaging Binaries...
	@mkdir -p tmp/$(PKG_NAME)
	@cp $(PLUGIN_PKG_NAME) tmp/$(PKG_NAME)/
	@cp extras/docs/* tmp/$(PKG_NAME)/
	@mkdir -p $(DIR)/dist/
	tar -czf $(DIR)/dist/$@ -C tmp $(PKG_NAME)
	@rm -f $(PKG_NAME) $(PLUGIN_PKG_NAME)
	@rm -rf tmp

clean:
	rm -f $(PKG_NAME) $(PLUGIN_PKG_NAME)
	rm -rf dist

.PHONY: dist all clean darwin_amd64_dist windows_amd64_dist linux_amd64_dist \
	install docker-release release pxc test

