# 默认执行 all 目标
.DEFAULT_GOAL := all

# ==============================================================================
# 定义 Makefile all 伪目标，执行 `make` 时，会默认会执行 all 伪目标
.PHONY: all
all: gen.add-copyright go.format go.lint go.build

# ==============================================================================
# Includes

# 确保 `include common.mk` 位于第一行，common.mk 中定义了一些变量，后面的子 makefile 有依赖
include scripts/make-rules/common.mk
include scripts/make-rules/tools.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/generate.mk

# ==============================================================================
# Usage

define USAGE_OPTIONS

Options:
  BINS             The binaries to build. Default is all of cmd.
                   This option is available when using: make build/build.multiarch
                   Example: make build BINS="miniblog test"
  VERSION          The version information compiled into binaries.
                   The default is obtained from gsemver or git.
  V                Set to 1 enable verbose build. Default is 0.
  tools			   tools.verify for verify tools
   				   tools.install for install tools
endef
export USAGE_OPTIONS

## --------------------------------------
## Generate / Manifests
## --------------------------------------

##@ generate:

.PHONY: add-copyright
add-copyright:
	@$(MAKE) gen.add-copyright

.PHONY: ca
ca:
	@$(MAKE) gen.ca

.PHONY: protoc
protoc:
	@$(MAKE) gen.protoc

.PHONY: deps
deps:
	@$(MAKE) gen.deps

## --------------------------------------
## Binaries
## --------------------------------------

##@ build:

.PHONY: build
build:
	@$(MAKE) go.build

## --------------------------------------
## Cleanup
## --------------------------------------

##@ clean:

.PHONY: clean
clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)


## --------------------------------------
## Lint / Verification
## --------------------------------------

##@ lint and verify:

.PHONY: lint
lint:
	@$(MAKE) go.lint


## --------------------------------------
## Hack / Tools
## --------------------------------------

##@ hack/tools:

.PHONY: format
format:
	@$(MAKE) go.format

.PHONY: swagger
swagger:
	@swagger serve -F=swagger --no-open --port 65534 $(ROOT_DIR)/api/openapi/openapi.yaml

.PHONY: tidy
tidy:
	@$(MAKE) go.tidy

.PHONY: help
help: Makefile
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<TARGETS> <OPTIONS>\033[0m\n\n\033[35mTargets:\033[0m\n"} /^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' Makefile #$(MAKEFILE_LIST)
	@echo -e "$$USAGE_OPTIONS"