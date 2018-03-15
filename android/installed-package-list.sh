#!/bin/bash
sdkmanager --list --verbose 2>/dev/null | sed -e '/Installed packages:/d' | sed -n '/Available Packages/!p;//q' | grep '^[a-z].*' | sed -e '/tools/d' >packages.txt
