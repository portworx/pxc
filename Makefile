CLINAME := px
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
#VER := $(shell git describe --tags)
VER := 0.0.0-$(SHA)
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
LDFLAGS :=-ldflags "-X github.com/portworx/px/cmd.PxVersion=$(VERSION)"

ifneq (windows,$(GOOS))
PKG_NAME = $(CLINAME)
else
PKG_NAME = $(CLINAME).exe
endif

PACKAGE := $(CLINAME)-$(VERSION).$(GOOS).$(ARCH).zip

all: px

install:
	go install

px:
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

$(PACKAGE): all
	@echo Packaging client Binaries...
	@mkdir -p dist
	@zip dist/$@ $(PKG_NAME)
	@rm -f $(PKG_NAME)

clean:
	rm -f $(PKG_NAME)
	rm -rf dist

.PHONY: dist all clean darwin_amd64_dist windows_amd64_dist linux_amd64_dist \
	install release px test

