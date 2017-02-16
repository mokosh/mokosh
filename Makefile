# The import path is where your repository can be found.
# To import subpackages, always prepend the full import path.
# If you change this, run `make clean`. Read more: https://git.io/vM7zV
IMPORT_PATH := github.com/mokosh/mokosh

V := 1 # When V is set, print commands and build progress.

allpackages = ./cmd/... ./pkg/...
allpackages_comma = "./cmd/...,./pkg/..."

# Space separated patterns of packages to skip in list, test, format.
IGNORED_PACKAGES := /vendor/

.PHONY: all
all: mokosrv

#.PHONY: build
#build: .GOPATH/.ok
#	$Q go install $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)

.PHONY: mokosrv
mokosrv: .GOPATH/.ok
	$Q go install $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)/cmd/mokoshsrv

##### ^^^^^^ EDIT ABOVE ^^^^^^ #####

##### =====> Utility targets <===== #####

.PHONY: clean test list cover format

clean:
	$Q rm -rf bin .GOPATH

test: .GOPATH/.ok
	$Q go test $(if $V,-v) -i -race $(allpackages)
ifndef CI
	$Q go vet $(allpackages)
ifeq ($(OS),Windows_NT)
	$Q go test -race $(allpackages)
else
	$Q GODEBUG=cgocheck=2 go test -race $(allpackages)
endif
else
	$Q ( go vet $(allpackages); echo $$? ) | \
	    tee .GOPATH/test/vet.txt | sed '$$ d'; exit $$(tail -1 .GOPATH/test/vet.txt)
	$Q ( GODEBUG=cgocheck=2 go test -v -race $(allpackages); echo $$? ) | \
	    tee .GOPATH/test/output.txt | sed '$$ d'; exit $$(tail -1 .GOPATH/test/output.txt)
endif

list: .GOPATH/.ok
	@echo $(allpackages)

cover: .GOPATH/.ok
	@echo "NOTE: make cover does not exit 1 on failure, don't use it to check for tests success!"
ifeq ($(OS),Windows_NT)
	$Q powershell if (Test-Path .GOPATH\cover) {Remove-Item .GOPATH\cover -Force -Recurse }
	$Q mkdir .GOPATH\cover
	go test -coverpkg=$(allpackages_comma) -coverprofile=.GOPATH\cover\unit-cmd.out ./cmd/...
	go test -coverpkg=$(allpackages_comma) -coverprofile=.GOPATH\cover\unit-pkg.out ./pkg/...
	$Q gocovmerge .GOPATH\cover\unit-cmd.out > .GOPATH\cover\all.merged
	$Q gocovmerge .GOPATH\cover\unit-pkg.out > .GOPATH\cover\all.merged
else
	$Q rm -f .GOPATH/cover/*.out .GOPATH/cover/all.merged
	$(if $V,@echo "-- go test -coverpkg=./... -coverprofile=.GOPATH/cover/... ./...")
	@for MOD in $(allpackages); do \
		go test -coverpkg=`echo $(allpackages)|tr " " ","` \
			-coverprofile=.GOPATH/cover/unit-`echo $$MOD|tr "/" "_"`.out \
			$$MOD 2>&1 | grep -v "no packages being tested depend on"; \
	done
	$Q gocovmerge .GOPATH/cover/*.out > .GOPATH/cover/all.merged
endif
ifndef CI
	$Q go tool cover -html .GOPATH/cover/all.merged
else
	$Q go tool cover -html .GOPATH/cover/all.merged -o .GOPATH/cover/all.html
endif
	@echo ""
	@echo "=====> Total test coverage: <====="
	@echo ""
	$Q go tool cover -func .GOPATH/cover/all.merged

format: .GOPATH/.ok
	$(foreach pkg,$(allpackages),$($Q go fmt $(pkg)))

##### =====> Internals <===== #####
VERSION          := $(shell git describe --tags --always --dirty="-dev")
ifeq ($(OS),Windows_NT)
DATE             := $(shell powershell get-date -format "{yyyy-mm-dd HH:mm:ss}")
else
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')
endif
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'

unexport GOBIN

Q := $(if $V,,@)

.GOPATH/.ok:
ifeq ($(OS),Windows_NT)
	$Q ECHO pie> $@
else
	$Q mkdir -p .GOPATH/test
	$Q mkdir -p .GOPATH/cover
	$Q touch $@
endif




# Based on https://github.com/cloudflare/hellogopher - v1.1 - MIT License
#
# Copyright (c) 2017 Cloudflare
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