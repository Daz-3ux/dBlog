# ==============================================================================
# define global Makefile variables for later reference

SHELL := /bin/bash

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/../../ && pwd -P))
OUTPUT_DIR := $(ROOT_DIR)/_output

# package name
ROOT_PACKAGE=github.com/Daz-3ux/dBlog

# Protobuf file path
APIROOT=$(ROOT_DIR)/pkg/proto

# ==============================================================================
# define version-related variables

# specify the version package used by the application
# values will be injected into the package using '-ldflags -X'
VERSION_PACKAGE=github.com/Daz-3ux/dBlog/pkg/version

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


# support OS list: linux/windows/darwin
PLATFORMS ?= darwin_amd64 windows_amd64 linux_amd64 linux_arm64

# species a OS
ifeq ($(origin PLATFORM), undefined)
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS)
	endif
	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
endif

# Makefile settings
ifndef V
MAKEFLAGS += --no-print-directory
endif

# Linux command settings
FIND := find . ! -path './third_party/*' ! -path './vendor/*'
XARGS := xargs --no-run-if-empty