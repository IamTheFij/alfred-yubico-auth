#! /bin/bash
set -e

echo "Warning! This will remove the workflow at the provided path and replace it with a link to this directory"
read -p "Path to workflow to replace: " existing_workflow

rm -fr "$existing_workflow"
ln -s `pwd` "$existing_workflow"
