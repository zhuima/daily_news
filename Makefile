PACKAGE  = News
DATE    ?= $(shell date +%FT%T%z)
VERSION  = 0.1.1
AUTHOR   = zhuima


.PHONY: all
all: ; $(info $(M)building executable...)
	$Q  go build \
		-tags release \
		-ldflags '-w -s -X main.GitTag=$(VERSION) -X main.BuildTime=$(DATE)' \
		-o bin/news main.go


.PHONY: clean
clean: ; $(info $(M) cleaning...)
	@rm -rf bin


.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo 'Version': $(VERSION)
	@echo 'Author': $(AUTHOR)
