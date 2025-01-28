#!/bin/bash

# คอมไพล์สำหรับ Linux
GOOS=linux GOARCH=amd64 go build -o app-for-linux ../main.go

# คอมไพล์สำหรับ Windows
GOOS=windows GOARCH=amd64 go build -o app-for-windows.exe ../main.go

# คอมไพล์สำหรับ macOS
GOOS=darwin GOARCH=amd64 go build -o app-for-macos ../main.go