#! /bin/bash

echo "$1" > ./alfred_yauth/version
plutil -replace version -string "$1" ./info.plist
