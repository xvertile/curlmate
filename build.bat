@echo off
set GO_FILE=curlmate.go
set OUTPUT_NAME=curlmate

echo Building for Windows...
set GOOS=windows
set GOARCH=amd64
go build -o windows\%OUTPUT_NAME%.exe %GO_FILE%

echo Building for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o linux\%OUTPUT_NAME% %GO_FILE%

echo Building for macOS...
set GOOS=darwin
set GOARCH=amd64
go build -o macos\%OUTPUT_NAME% %GO_FILE%

echo Build complete!
