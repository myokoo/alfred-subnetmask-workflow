#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0);pwd)

LIB1="$HOME/Library/Application Support/Alfred 3/Alfred.alfredpreferences"
LIB2="$HOME/Library/Application Support/Alfred/Alfred.alfredpreferences"

[[ -d "$LIB1" ]] && LIB="$LIB1/workflows"
[[ -d "$LIB2" ]] && LIB="$LIB2/workflows"

[[ ! -d "$LIB" ]] && mkdir $LIB
cd "$LIB" || exit 2

ln -fs "$SCRIPT_DIR" "$LIB/SubnetMask" && echo "setup ok,"
