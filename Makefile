all: version

VERSION_TAG := `date -u "+%Y%m.%d-build"`
MAIN_VERSION := `./go/cmd/goProperties -file='$(FILENAME)' -key='major'`
version:
	./go/cmd/goProperties -file='$(FILENAME)' -key='version' -value="$(MAIN_VERSION).$(VERSION_TAG).$(GITHUB_RUN_ID)" -save

#TODO set the ROOT_DIR var
ROOT_DIR := src\/avro\/schm\/
DIR_SEP := \/
REPLACE_PATH := ${ROOT_DIR}${DOMAIN}${DIR_SEP}${SUBDOMAIN}
CHECKDIR := ./$(shell echo ${ROOT_DIR} | sed -e "s/\\\//g")${DOMAIN}/${SUBDOMAIN}
create:
	if [ -z "${DOMAIN}" ]; then echo "::error ::DOMAIN is required" && exit 1; fi
	if [ -z "${SUBDOMAIN}" ]; then echo "::error ::SUBDOMAIN is required" && exit 1; fi
	if [ -d "${CHECKDIR}" ]; then echo "::error ::DIR ${CHECKDIR} already exists" && exit 1; fi
	mkdir -p ${CHECKDIR}
	cp templates/version.properties ${CHECKDIR}/version.properties
	cat templates/subdomain.workflow.yml | \
		sed -e "s/{{__DOMAIN__}}/${DOMAIN}/" | \
		sed -e "s/{{__SUBDOMAIN__}}/${SUBDOMAIN}/" | \
		sed -e "s/{{__PATH__}}/${REPLACE_PATH}/" | \
		sed -e "s/{{__NAME__}}/${DOMAIN}.${SUBDOMAIN}/" > .github/workflows/$(DOMAIN).$(SUBDOMAIN).yml

jar:
	if [ -z "${DOMAIN}" ]; then echo "::error ::DOMAIN is required" && exit 1; fi
	if [ -z "${SUBDOMAIN}" ]; then echo "::error ::SUBDOMAIN is required" && exit 1; fi
	if [ -z "${VERSION}" ]; then echo "::error ::VERSION is required" && exit 1; fi
	rm -Rf build targets
	JAR_VERSION=${VERSION} DOMAIN=${DOMAIN} SUBDOMAIN=${SUBDOMAIN} ./gradlew clean build

.PHONY: version all create jar