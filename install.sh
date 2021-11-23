#!/usr/bin/env bash

# Install the Vamp Cloud CLI tool.
# https://github.com/magneticio/vamp-cloud-cli
#
# Dependencies: curl, grep, sed
#
# The version to install and the binary location can be passed in via VERSION and DESTDIR respectively.
#

set -o errexit

echo "Starting installation."

# GitHub's URL for the latest release, will redirect.
GITHUB_BASE_URL="https://github.com/magneticio/vamp-cloud-cli"
GITHUB_API_URL="https://api.github.com/repos/magneticio/vamp-cloud-cli/releases/latest"
DESTDIR="${DESTDIR:-/usr/local/bin}"

if [ -z "$VERSION" ]; then
  VERSION=$(curl -s $GITHUB_API_URL | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
fi

echo "Installing Vamp Cloud CLI ${VERSION}"

trap error ERR

# Determine release filename. This can be expanded with CPU arch in the future.
case "$(uname)" in
  Linux)
    OS='linux'
    ARCH="$(uname -m)"
  ;;
  Darwin)
    OS='darwin'
    ARCH="$(uname -m)"
    if [ "$ARCH" = "x86_64" ]; then
      if [ "$(sysctl -in sysctl.proc_translated)" = "1" ]; then
          # Running on Rosetta 2"
          ARCH='arm64'
      fi
    fi
  ;;
  *)
    echo "$OS: This operating system is not supported."
    echo "You may be able to manually download the CLI at"
    echo "https://github.com/magneticio/vamp-cloud-cli/releases/latest"
    exit 1  
  ;;
esac

RELEASE_URL="${GITHUB_BASE_URL}/releases/download/${VERSION}/vamp-${OS}-${ARCH}"

# Download & add execute perm
curl -sL --retry 3 "${RELEASE_URL}" -o "$DESTDIR/vamp"
chmod +x "$DESTDIR/vamp"

echo "Installed to $DESTDIR"

command -v vamp

vamp version