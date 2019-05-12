#!/bin/sh

set -eu

RELEASE="${RELEASE:-"1.0.0"}"

URL="https://github.com/piotrpersona/httpst/releases/download/${RELEASE}/httpst.tar.gz"
ARCHIVE_PATH="/tmp/httpst.tar.gz"
PACKAGE_PATH="/var/lib/httpst"
HTTPST_BIN="/usr/local/bin/httpst"

echo "INFO: Fetching release ${RELEASE} "
curl -fSL "${URL}" -o "${ARCHIVE_PATH}"

echo "INFO: Unpacking archive to ${PACKAGE_PATH}"
mkdir -p "${PACKAGE_PATH}"
tar -C "${PACKAGE_PATH}" -xzf "${ARCHIVE_PATH}"
mv "${PACKAGE_PATH}/tmp/httpst"/* "${PACKAGE_PATH}" \
    && rm -rf "${PACKAGE_PATH}/tmp"

echo "INFO: Linking ${PACKAGE_PATH}/httpst to ${HTTPST_BIN}"
ln -f "${PACKAGE_PATH}/httpst" "${HTTPST_BIN}"

rm -rf "${ARCHIVE_PATH}"

echo "INFO: Done."
