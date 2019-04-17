# TEMPLATE MAKEFILE FOR GO APPLICATIONS
# 
# This Makefile is capable of:
#  - defautl go commands like build, run,
#    test and golint
#  - cross compiling (see make help)
#  - automatic versioning using git tag
#    and ldflags
#
# Covered by MIT License
#
# Copyright 2019 Ringo Hoffmann (zekro Development)
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

### NAMES AND LOCS ############################
APPNAME   = thunder
PACKAGE   = github.com/zekroTJA/$(APPNAME)
EXAMPLE   = $(CURDIR)/examples/userdb.go
TESTPKGS  = ./...
# golint min_confidence: see'golint --help'
LINTCONF  = 0
###############################################

### EXECUTABLES ###############################
GO     = go
GOLINT = golint
###############################################

# ---------------------------------------------

# TAG    = $(shell git describe --tags)
# COMMIT = $(shell git rev-parse HEAD)


ifneq ($(GOOS),)
	BIN := $(BIN)_$(GOOS)
endif

ifneq ($(GOARCH),)
	BIN := $(BIN)_$(GOARCH)
endif

ifneq ($(TAG),)
	BIN := $(BIN)_$(TAG)
endif

ifeq ($(OS),Windows_NT)
	BIN := $(BIN).exe
endif


PHONY = _make
_make: cleanup tests

PHONY += tests
tests:
	$(GO) test $(TESTPKGS) -v -cover

PHONY += test
test: tests

PHONY += lint 
lint: 
	$(GOLINT) -min_confidence $(LINTCONF) ./...

PHONY += example
example:
	$(GO) run -v $(EXAMPLE)

PHONY += cleanup
cleanup:
	rm -f $(CURDIR)/tests/*.th
	rm -f $(CURDIR)/examples/*.th
	rm -f $(CURDIR)/*.th

PHONY += help
help:
	@echo "Available targets:"
	@echo "  #        - creates binary in ./bin"
	@echo "  cleanup  - tidy up temporary stuff created by tests or examples"
	@echo "  example  - run example"
	@echo "  lint     - run linters (golint)"
	@echo "  test     - run tests (go test)"
	@echo ""


.PHONY: $(PHONY)