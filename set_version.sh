#! /bin/bash

plutil -replace version -string "$1" ./info.plist
