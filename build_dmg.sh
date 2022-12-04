#!/bin/bash

# The script in this file modified the script downloaded from the following source:
# https://gist.github.com/blockpane/fe03eb0839fac417b92cd7eb98cdf356

# builds a fyne macos app, signs it, creates a .dmg, signs that, and then requests notarization.
# ... stapling the notarization should wait until the request is approved which can take a minute ...
#
# Pre-reqs:
#   Install the full xcode package
#   Install the fyne command line tool
#   Create a certificate request, send to apple developer site, and import the resulting certificate
#   Create an application password in appleid.apple.com tied to the app's name in the keychain

# Go to https://developer.apple.com/account/resources/certificates/list
# create a "Developer ID Application" certificate
# Go to https://developer.apple.com/account/resources/identifiers/list
# select "Register a new identifier" > "App IDs" > continue
# save the name and ID of the app

# The name of the application
APP_NAME="GoBible"
# The bundle / app ID for the app...
APP_ID="app.gobible"
# This is the CN from the code signing cert
CERT="Developer ID Application: XXXXX (XXXXX)"
# This is your login for the apple developer site
EMAIL="xxx@xxx.com"
# To set up a password for the app
# Go to https://appleid.apple.com/
# Select "App-specific passwords"
# generate
# give it a name which is the same name as APP_ID
# enter Apple ID password
# write down the generated password, appeared on the screen
# enter the pass below
KEYCHAIN_PASS="xxxx-xxxx-xxxx-xxxx"
# A working dir where temp files are created, and final disk image will be placed.
PKG_DIR="dmg"
# The entrypoint to compile:
MAIN="gobible.go"
# Set to "TRUE" if the app version should use the git tag:
TAG="TRUE"
### End of settings ###

# should permit a package built on MacOS versions >= 10.15 to run on 10.14 and later, maybe.
export CGO_CFLAGS="-mmacosx-version-min=10.14"
export CGO_LDFLAGS="-mmacosx-version-min=10.14"

set -x
# don't overwrite something that already exists when we compile our bin.
BIN=$(dirname "${MAIN}" |awk -F/ '{print $NF}')
MAIN_DIR=$(dirname "${MAIN}")
[ -f "${MAIN_DIR}/${BIN}" ] && { echo "refusing to continue, ${MAIN_DIR}/${BIN} already exists"; kill 0; }

# Cleanup old packages if needed
[ -d "${PKG_DIR}" ] && { mkdir -p "${PKG_DIR}/old"; mv "${PKG_DIR}/"*.dmg "${PKG_DIR}/old/"; }
[ -d "${PKG_DIR}/${APP_NAME}" ] && rm -fr "${PKG_DIR}/${APP_NAME}"
mkdir -p "${PKG_DIR}/${APP_NAME}"

go build -ldflags "-s -w" -o "${MAIN_DIR}/${BIN}" "${MAIN}" || exit

fyne package -sourceDir "${MAIN_DIR}" -name "${APP_NAME}" -os darwin -appID "${APP_ID}" && mv "${APP_NAME}.app" "${PKG_DIR}/${APP_NAME}/" || exit
rm -f "${MAIN_DIR}/${BIN}"

[ "${TAG}" == "TRUE" ] && sed -i'' -e 's/.string.1\.0.\/string./\<string>'$(git describe --tags --always --long)'\<\/string>/g' "${PKG_DIR}/${APP_NAME}/${APP_NAME}.app/Contents/Info.plist"

pushd "${PKG_DIR}"; pushd "${APP_NAME}"

# remove any existing extended attributes from new app:
xattr -rc "${APP_NAME}.app"
# Add a link for /Applications in the Directory, more covenient when becomes a .dmg:
ln -s /Applications Applications

# sign the app:
codesign --force --options runtime --deep --sign "${CERT}" -i "${APP_ID}" "${APP_NAME}.app" || exit

popd
# Create a disk image
hdiutil create -srcfolder "${APP_NAME}" "${APP_NAME}.dmg" || exit
# sign the .dmg
codesign --sign "${CERT}" -i "${APP_ID}" "${APP_NAME}.dmg" || exit

# export the PATH of altool
export PATH=$PATH:/Applications/Xcode.app/Contents/Developer/usr/bin

# Warning: altool has been deprecated for notarization and starting in late 2023 will no longer be supported by the Apple notary service. You should start using notarytool to notarize your software.
# Request anotization from apple
UUID=$(xcrun altool -t osx -f "${APP_NAME}".dmg --primary-bundle-id "${APP_ID}" --notarize-app --username "${EMAIL}" -p "${KEYCHAIN_PASS}" | grep RequestUUID | awk '{print $NF}')
rm -fr "${APP_NAME}/"

set +x
# wait until approved and then staple the dmg
while true; do
        sleep 5
        echo checking....
        xcrun altool --notarization-info "${UUID}" -p "${KEYCHAIN_PASS}" --username "${EMAIL}" | grep -s Approved && break
done
xcrun stapler staple -v "${APP_NAME}.dmg"
