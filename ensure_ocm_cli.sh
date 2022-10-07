#!/bin/bash
set -eu
ocmDownloads=$(curl -s https://api.github.com/repos/openshift-online/ocm-cli/releases/latest | grep "browser_download_url.*" | cut -d : -f 2,3 | tr -d \")
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
mkdir tests/ocm
wget -P tests/ocm -q $downloadURL
mv tests/ocm/ocm-$osName-$hardware tests/ocm/ocm
