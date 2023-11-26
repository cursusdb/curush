# Curush
The CursusDB Shell

https://cursusdb.com

## Using
X = host such as an IP or FQDN
``` 
/curush --host=X
```

You can specify a specific CursusDB cluster port.
``` 
/curush --host=X --port=1234
```

## Building

### Darwin / MacOS

- ``env GOOS=darwin GOARCH=arm go build -o bin/macos-darwin/arm/curush``

- ``env GOOS=darwin GOARCH=amd64 go build -o bin/macos-darwin/amd64/curush``

- ``env GOOS=darwin GOARCH=386 go build -o bin/macos-darwin/386/curush``

- ``env GOOS=darwin GOARCH=arm64 go build -o bin/macos-darwin/arm64/curush``

### Linux
- ``env GOOS=linux GOARCH=386 go build -o bin/linux/386/curush``

- ``env GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/curush``

- ``env GOOS=linux GOARCH=arm go build -o bin/linux/arm/curush``

``env GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/curush``

### FreeBSD

- ``env GOOS=freebsd GOARCH=arm go build -o bin/freebsd/arm/curush``

- ``env GOOS=freebsd GOARCH=amd64 go build -o bin/freebsd/amd64/curush``

- ``env GOOS=freebsd GOARCH=386 go build -o bin/freebsd/386/curush``


### Windows
- ``env GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/curush``
- ``env GOOS=windows GOARCH=386 go build -o bin/windows/386/curush``

