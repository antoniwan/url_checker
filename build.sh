#!/bin/bash

# Compile for Windows
GOOS=windows GOARCH=amd64 go build -o url_checker.exe url_checker.go

# Compile for Mac
GOOS=darwin GOARCH=amd64 go build -o url_checker_mac url_checker.go

echo "Compilation completed."