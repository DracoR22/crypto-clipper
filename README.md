<div>
  <h1 align="center">
    Crypto Clipper 🪙
  </h1>
</div>

<p align="center">
 Built with GOLANG. Supports Bitcoin, Ethereum, Solana, Cosmos
</p>

## Build the code for your current OS:
```shell
go build -ldflags -H=windowsgui -o main.exe
```

## Build Cross-Platform:

### Install Go Releaser
```shell
go install github.com/goreleaser/goreleaser/v2@latest
```

### Build the cross platform code
```shell
goreleaser --snapshot --clean                          
```