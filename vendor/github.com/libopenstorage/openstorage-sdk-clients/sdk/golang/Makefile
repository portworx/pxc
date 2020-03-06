HAS_PROTOC_GEN_GO := $(shell command -v protoc-gen-go 2> /dev/null)
APIDIR := ../..
ifndef PROTOC
PROTOC = protoc
endif

ifndef PROTOS_PATH
PROTOS_PATH = $(GOPATH)/src
endif

ifndef PROTOSRC_PATH
PROTOSRC_PATH = $(APIDIR)
endif

all: grpc

clean:

grpc:
ifndef HAS_PROTOC_GEN_GO
	$(error "Please install grpc rools for ruby: gem install grpc && gem install grpc-tools")
endif
	$(PROTOC) -I $(PROTOSRC_PATH) \
		-I /usr/local/include \
		-I $(PROTOS_PATH)/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		$(PROTOSRC_PATH)/api.proto
	$(PROTOC) -I $(PROTOSRC_PATH) \
		-I /usr/local/include \
		-I $(PROTOS_PATH)/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		$(PROTOSRC_PATH)/api.proto

.PHONY: all grpc
