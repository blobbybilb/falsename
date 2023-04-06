# build for all platforms, and save to ./build/<platform>/fn

# build for linux
GOOS=linux GOARCH=amd64 go build -o build/linux-amd64/fn
GOOS=linux GOARCH=arm64 go build -o build/linux-arm64/fn

# build for mac
GOOS=darwin GOARCH=amd64 go build -o build/mac-amd64/fn
GOOS=darwin GOARCH=arm64 go build -o build/mac-arm64/fn

# build for windows
GOOS=windows GOARCH=amd64 go build -o build/windows-amd64/fn.exe

# build for freebsd
GOOS=freebsd GOARCH=amd64 go build -o build/freebsd-amd64/fn
GOOS=freebsd GOARCH=arm64 go build -o build/freebsd-arm64/fn