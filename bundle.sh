#!/bin/bash
VERSION=v1.1.0

echo "Bundling shell zip and tars into /bin in current directory."

env GOOS=darwin GOARCH=amd64 go build -o bin/macos-darwin/amd64/curush && tar -czf bin/macos-darwin/amd64/curush-$VERSION-amd64.tar.gz -C bin/macos-darwin/amd64/ $(ls  bin/macos-darwin/amd64/)
env GOOS=darwin GOARCH=arm64 go build -o bin/macos-darwin/arm64/curush && tar -czf bin/macos-darwin/arm64/curush-$VERSION-arm64.tar.gz -C bin/macos-darwin/arm64/ $(ls  bin/macos-darwin/arm64/)
env GOOS=linux GOARCH=386 go build -o bin/linux/386/curush && tar -czf bin/linux/386/curush-$VERSION-386.tar.gz -C bin/linux/386/ $(ls  bin/linux/386/)
env GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/curush && tar -czf bin/linux/amd64/curush-$VERSION-amd64.tar.gz -C bin/linux/amd64/ $(ls  bin/linux/amd64/)
env GOOS=linux GOARCH=arm go build -o bin/linux/arm/curush && tar -czf bin/linux/arm/curush-$VERSION-arm.tar.gz -C bin/linux/arm/ $(ls  bin/linux/arm/)
env GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/curush && tar -czf bin/linux/arm64/curush-$VERSION-arm64.tar.gz -C bin/linux/arm64/ $(ls  bin/linux/arm64/)
env GOOS=freebsd GOARCH=arm go build -o bin/freebsd/arm/curush && tar -czf bin/freebsd/arm/curush-$VERSION-arm.tar.gz -C bin/freebsd/arm/ $(ls  bin/freebsd/arm/)
env GOOS=freebsd GOARCH=amd64 go build -o bin/freebsd/amd64/curush && tar -czf bin/freebsd/amd64/curush-$VERSION-amd64.tar.gz -C bin/freebsd/amd64/ $(ls  bin/freebsd/amd64/)
env GOOS=freebsd GOARCH=386 go build -o bin/freebsd/386/curush && tar -czf bin/freebsd/386/curush-$VERSION-386.tar.gz -C bin/freebsd/386/ $(ls  bin/freebsd/386/)
env GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/curush.exe && zip -r -j bin/windows/amd64/curush-$VERSION-x64.zip bin/windows/amd64/curush.exe
env GOOS=windows GOARCH=arm64 go build -o bin/windows/arm64/curush.exe && zip -r -j bin/windows/arm64/curush-$VERSION-x64.zip bin/windows/arm64/curush.exe
env GOOS=windows GOARCH=386 go build -o bin/windows/386/curush.exe && zip -r -j bin/windows/386/curush-$VERSION-x86.zip bin/windows/386/curush.exe

echo "Fin"