all: version

#VERSION_TAG := `date -u "+%Y%m.%d-build.%H%M%S"`
VERSION_TAG := `date -u "+%Y%m.%d-build"`
MAIN_VERSION := `./go/cmd/goProperties-mac -file='$(FILENAME)' -key='major'`

version:
	./go/cmd/goProperties-mac -file='$(FILENAME)' -key='version' -value="$(MAIN_VERSION).$(VERSION_TAG).$(GITHUB_RUN_ID)" -save

.PHONY: version all