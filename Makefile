# ==============================================================================
# define global Makefile variables for later reference

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# project root directory
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))
# directory for storing build output and temporary files
OUTPUT_DIR := $(ROOT_DIR)/_output


# ==============================================================================
# define version-related variables

# specify the version package used by the application
# values will be injected into the package using '-ldflags -X'
VERSION_PACKAGE=github.com/daz-3ux/dBlog/pkg/version

## define readable version number
ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

## check if the code repository is dirty(default is dirty)
GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif
GIT_COMMIT:=$(shell git rev-parse HEAD)

GO_LDFLAGS += \
	-X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

# ==============================================================================
# define the Makefile "all" phony target, which is executed by default when running 'make'
.PHONY: all
all: add-copyright format build

# ==============================================================================
# define other phony targets

.PHONY: build
build: tidy # compile source code, auto adding/removing dependency packages depending on "tidy" target
	go build -v -ldflags "$(GO_LDFLAGS)" -o $(OUTPUT_DIR)/dBlog $(ROOT_DIR)/cmd/dBlog/main.go

.PHONY: format
format: # format source code
	gofmt -s -w ./

.PHONY: add-copyright
add-copyright: # add license header
	addlicense -v -f $(ROOT_DIR)/scripts/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,$(OUTPUT_DIR)

.PHONY: swagger
swagger: # start swagger ui
	swagger serve -F=swagger --no-open --port 65534 $(ROOT_DIR)/api/openapi/openapi.yaml

.PHONY: tidy
tidy: # auto add/remove dependency packages
	go mod tidy

.PHONY: clean
clean: # clean build output and temporary files
	-rm -vrf $(OUTPUT_DIR)
