#!/bin/bash
if which sdkmanager >/dev/null; then
    sdkmanager --list --verbose 2>/dev/null | sed -e '/Installed packages:/d' | sed -n '/Available Packages/!p;//q' | grep '^[a-z].*' | sed -e '/tools/d' >packages.txt
else
    echo "Couldn't find sdkmanager in your PATH"
    exit 1
fi
