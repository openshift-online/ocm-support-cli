#!/bin/bash
# This script downloads latest binaries of OCM CLI and stores them under tests/ocm for managing OCM CLI specific test dependencies
set -eu
echo "Getting download URLs from GitHub"
ocmDownloads=$(curl -s https://api.github.com/repos/openshift-online/ocm-cli/releases/latest | grep "browser_download_url.*" | cut -d : -f 2,3 | tr -d \")
# currently supported OS : darwin-amd64, linux-amd64, linux-arm64, linux-ppc64le, linux-s390
echo "Checking system OS and hardware"
osName=$(uname)
osName="${osName,,}"
hardware=$(uname -m)
amd64="x86_64"
if [[ "$hardware" = "$amd64" ]]; then
    hardware="amd64"
fi
downloadURL=""
ocmCLI=ocm-$osName-$hardware
echo "Finding the URL for the matching system"
for addr in $ocmDownloads
do
    if [[ "$addr" = *$ocmCLI ]]; then
        downloadURL="$addr"
        break
    fi
done
echo "Recreating tests/ocm folder to download latest binaries"
rm -rf tests/ocm
mkdir -p tests/ocm
echo "Downloading latest binaries of OCM CLI"
wget -P tests/ocm -q $downloadURL
echo "Setting execute permission to binaries"
chmod +rwx tests/ocm/$ocmCLI
mv tests/ocm/$ocmCLI tests/ocm/ocm
echo "OCM CLI binaries installed successfully"