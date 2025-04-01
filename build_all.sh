#!/bin/bash

APP_NAME="mini-irc"
DIST_DIR="dist"

rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

platforms=("darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm64" "windows/amd64")

for platform in "${platforms[@]}"
do
    IFS="/" read -r GOOS GOARCH <<< "$platform"
    output_name="$APP_NAME-$GOOS-$GOARCH"
    if [ "$GOOS" == "windows" ]; then
        output_name="$output_name.exe"
    fi

    echo "Building for $GOOS/$GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o "$DIST_DIR/$output_name" main.go
done

