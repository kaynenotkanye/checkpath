#!/bin/bash
### This is a generic script to package up go as a compiled binary

rm ./bin/checkpath

echo "Building binary"
GOOS=linux GOARCH=amd64 go build -o checkpath checkpath.go

mv checkpath ./bin/
