all: version

VERSION_TAG := `date -u "+%Y.%m.%d-%H%M%S"`

version:
	echo "version = $(VERSION_TAG)" > $(FILENAME)

.PHONY: version all