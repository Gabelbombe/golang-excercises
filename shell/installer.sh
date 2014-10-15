#!/bin/bash
#

## ROOT check
if [[ $EUID -ne 0 ]]; then
  echo "This script must be run as su" 1>&2 ; exit 1
fi

cd /tmp
wget 'https://storage.googleapis.com/golang/go1.3.3.darwin-amd64-osx10.8.pkg'
installer -pkg 'go1.3.3.darwin-amd64-osx10.8.pkg' -target /

export GOROOT="/usr/local/go"
export GOPATH="$HOME/go"