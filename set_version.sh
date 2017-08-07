#! /bin/bash

echo "$1" > ./src/version
plutil -replace version -string "$1" ./src/info.plist
