#!/bin/bash

VERSION_PROPERTIES=$(./go/cmd/goProperties -file='gradle.properties' -prop='version')

if [ -z "$VERSION_PROPERTIES" ]
then
      echo "* Not found version in properties file!"
      exit 1
else
      echo "* Tag version from properties file v$VERSION_PROPERTIES"
fi

VERSION_REPO=$(git tag --list | grep "v$VERSION_PROPERTIES")

if [ -z "$VERSION_REPO" ]
then
      echo "* Available tag v$VERSION_PROPERTIES"
else
      echo "* Tag v$VERSION_PROPERTIES already exists in repo!"
      exit 1
fi