# Should match the name of component
CLINAME := cm

PXC_COMPONENT := pxc-$(CLINAME)
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
VER := $(shell git describe --tags)
ARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)
DIR=.
PKGDIR=$(DIR)/tmp/$(PXC_COMPONENT)

ifdef APP_SUFFIX
  VERSION = $(VER)-$(subst /,-,$(APP_SUFFIX))
else
ifeq (master,$(BRANCH))
  VERSION = $(VER)
else
  VERSION = $(VER)-$(BRANCH)
endif
endif
LDFLAGS :=-ldflags "-X main.ComponentName=$(CLINAME) -X main.ComponentVersion=$(VERSION)"

ifneq (windows,$(GOOS))
PKG_NAME = $(CLINAME)
COMPONENT_PKG_NAME = $(PXC_COMPONENT)
else
PKG_NAME = $(CLINAME).exe
COMPONENT_PKG_NAME = $(PXC_COMPONENT).exe
endif

ZIPPACKAGE := $(PXC_COMPONENT)-$(VERSION).$(GOOS).$(ARCH).zip
TGZPACKAGE := $(PXC_COMPONENT)-$(VERSION).$(GOOS).$(ARCH).tar.gz

all: $(PKG_NAME) $(COMPONENT_PKG_NAME)

install: all
	cp $(PKG_NAME) $(HOME)/.pxc/bin/$(COMPONENT_PKG_NAME)

imports:
	goimports -w ./cmd
	goimports -w ./handler
	goimports -w ./pkg
	goimports -w *.go

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

$(PKG_NAME):
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
	./$(PKG_NAME) gendocs --output-dir=docs/usage

# Look at pxc ./hack directory to see how it creates environments to test
#test:
#	./hack/test.sh

verify: all
	go fmt $(go list ./... | grep -v vendor) | wc -l | grep 0
	go vet $(go list ./... | grep -v vendor)

$(COMPONENT_PKG_NAME): $(PKG_NAME)
	cp $(PKG_NAME) $(COMPONENT_PKG_NAME)

$(ZIPPACKAGE): all
	@echo Packaging binaries...
	@mkdir -p $(PKGDIR)/bin
	@cp $(COMPONENT_PKG_NAME) $(PKGDIR)/bin
	@mkdir -p dist
	( cd tmp && zip -r ../dist/$@ $(PXC_COMPONENT) )
	@rm -f $(PKG_NAME)
	@rm -rf tmp

$(TGZPACKAGE): all
	@echo Packaging binaries...
	@mkdir -p $(PKGDIR)/bin
	@cp $(COMPONENT_PKG_NAME) $(PKGDIR)/bin
	@mkdir -p $(DIR)/dist/
	tar -czf $(DIR)/dist/$@ -C tmp $(PXC_COMPONENT)
	@rm -rf tmp

clean:
	rm -f $(PKG_NAME) $(COMPONENT_PKG_NAME) $(CLINAME).exe $(PXC_COMPONENT).exe
	rm -rf dist

.PHONY: dist all clean darwin_amd64_dist windows_amd64_dist linux_amd64_dist \
	install release test

