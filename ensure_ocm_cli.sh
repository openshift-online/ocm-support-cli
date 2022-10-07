#!/bin/bash
# This script downloads latest binaries of OCM CLI and stores them under tests/ocm for managing OCM CLI specific test dependencies
set -eu
ocmDownloads=$(curl -s https://api.github.com/repos/openshift-online/ocm-cli/releases/latest | grep "browser_download_url.*" | cut -d : -f 2,3 | tr -d \")
# currently supported OS : darwin-amd64, linux-amd64, linux-arm64, linux-ppc64le, linux-s390, windows-amd64
osName=$(uname)
osName="${osName,,}"
hardware=$(uname -m)
amd64="x86_64"
if [[ "$hardware" = "$amd64" ]]; then
    hardware="amd64"
fi
downloadURL=""
for addr in $ocmDownloads
do
    if [[ "$addr" = *ocm-$osName-$hardware ]]; then
        downloadURL="$addr"
        break
    fi
done
rm -rf tests/ocm
mkdir -p tests/ocm
wget -P tests/ocm -q $downloadURL
mv tests/ocm/ocm-$osName-$hardware tests/ocm/ocm
