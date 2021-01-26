# Should match the name of component
CLINAME := cm

PXC_COMPONENT := pxc-$(CLINAME)
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
VER := $(shell git describe --tags)
ARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)
DIR=.
TMP=$(DIR)/tmp
PKGDIR=$(TMP)/$(PXC_COMPONENT)

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
COMPONENT_PKG_NAME = $(PXC_COMPONENT)
else
COMPONENT_PKG_NAME = $(PXC_COMPONENT).exe
endif

ZIPPACKAGE := $(PXC_COMPONENT)-$(VERSION).$(GOOS).$(ARCH).zip
TGZPACKAGE := $(PXC_COMPONENT)-$(VERSION).$(GOOS).$(ARCH).tar.gz

all: component

install: all
	cp $(COMPONENT_PKG_NAME) $(GOPATH)/bin/$(COMPONENT_PKG_NAME)

imports:
	goimports -w ./cmd
	goimports -w ./handler
	goimports -w ./pkg
	goimports -w *.go

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

component:
	go build -o $(COMPONENT_PKG_NAME) $(LDFLAGS)

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
	./$(COMPONENT_PKG_NAME) gendocs --output-dir=docs/usage

# Look at pxc ./hack directory to see how it creates environments to test
#test:
#	./hack/test.sh

verify: all
	go fmt $(go list ./... | grep -v vendor) | wc -l | grep 0
	go vet $(go list ./... | grep -v vendor)

$(ZIPPACKAGE): all
	@echo Packaging binaries...
	@mkdir -p $(PKGDIR)/bin
	@cp $(COMPONENT_PKG_NAME) $(PKGDIR)/bin
	@mkdir -p dist
	( cd $(DIR)/tmp && zip -r ../dist/$@ $(PXC_COMPONENT) )
	@rm -rf $(TMP)
	@rm -f $(COMPONENT_PKG_NAME)

$(TGZPACKAGE): all
	@echo Packaging binaries...
	@mkdir -p $(PKGDIR)/bin
	@cp $(COMPONENT_PKG_NAME) $(PKGDIR)/bin
	@mkdir -p $(DIR)/dist/
	tar -czf $(DIR)/dist/$@ -C $(DIR)/tmp $(PXC_COMPONENT)
	@rm -rf $(TMP)
	@rm -f $(COMPONENT_PKG_NAME)

clean:
	rm -f $(COMPONENT_PKG_NAME)
	rm -rf dist

.PHONY: dist all clean darwin_amd64_dist windows_amd64_dist linux_amd64_dist \
	install release test component

