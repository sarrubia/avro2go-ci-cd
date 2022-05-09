all: version

VERSION_TAG := `date -u "+%Y.%m-%d_%H%M%S"`
MAIN_VERSION := `./go/cmd/goProperties -file='$(FILENAME)' -key='breaking'`

version:
	./go/cmd/goProperties -file='$(FILENAME)' -key='version' -value="$(MAIN_VERSION).$(VERSION_TAG)" -save

.PHONY: version all