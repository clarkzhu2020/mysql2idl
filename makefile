build:
    # 交叉编译为 Linux 系统
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mysql2idl -ldflags="-s -w" .
    # 交叉编译为 Windows 系统
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mysql2idl.exe -ldflags="-s -w" .
    # 交叉编译为 macOS 系统
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o mysql2idl -ldflags="-s -w" .

clean:
    rm -f main_linux_amd64 main_windows_amd64.exe main_darwin_amd64