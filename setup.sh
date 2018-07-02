#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0);pwd)

LIB="$HOME/Library/Application Support/Alfred 3/Alfred.alfredpreferences/workflows"
cd "$LIB" || exit 2

ln -fs "$SCRIPT_DIR" SubnetMask && echo "setup ok," || echo error
