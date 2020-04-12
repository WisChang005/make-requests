#! /bin/bash

TARGET_FOLDERS="../make-pageviews"

[ -d "${TARGET_FOLDERS}" ] && echo "Skip to create folder" || mkdir ../make-pageviews

echo "Copy config file..."
cp config.ini ${TARGET_FOLDERS}/.

echo "Build packages..."
go build -o ../make-pageviews/makePageviews.exe

echo "Build finished!"
