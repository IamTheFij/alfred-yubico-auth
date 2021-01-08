#! /bin/bash

# Statically compile the builder
mage -compile ./static-builder

# Change the environment to one that looks kinda like Alfred
export HOME=fake_alfred_home
alfred_sync_dir="$HOME/Library/Application Support/Alfred"
mkdir -p "$alfred_sync_dir/Alfred.alfredpreferences/workflows"
echo '{"syncfolders": {"4": "/dev/null"}, "current": "/dev/null"}' > "$alfred_sync_dir/prefs.json"
# Run the builder with whatever arguments
export GOOS=darwin
export GOARCH=amd64
./static-builder "$@"
