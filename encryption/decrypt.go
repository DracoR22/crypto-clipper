package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
	"os/exec"
)

func DecryptFile(encFilePath, decFilePath string, decryptionKey []byte) error {
	// Open encrypted file
	encFile, err := os.Open(encFilePath)
	if err != nil {
		return err
	}
	defer encFile.Close()

	// Create output file for decrypted data
	decFile, err := os.Create(decFilePath)
	if err != nil {
		return err
	}
	defer decFile.Close()

	// Read IV (Initialization Vector) from the beginning of the encrypted file
	iv := make([]byte, aes.BlockSize)
	if _, err := encFile.Read(iv); err != nil {
		return err
	}

	// Create AES cipher block
	block, err := aes.NewCipher(decryptionKey)
	if err != nil {
		return err
	}

	// Create CBC mode decrypter
	mode := cipher.NewCBCDecrypter(block, iv)

	// Buffer for reading encrypted file
	buffer := make([]byte, 4096)
	for {
		// Read a chunk of encrypted data
		bytesRead, err := encFile.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if bytesRead == 0 {
			break
		}

		// Decrypt the chunk
		decryptedData := make([]byte, bytesRead)
		mode.CryptBlocks(decryptedData, buffer[:bytesRead])

		// Write decrypted data to output file
		if _, err := decFile.Write(decryptedData); err != nil {
			return err
		}
	}

	return nil
}

func RunDecryptedExecutable(executablePath string) error {
	// Example: Run the decrypted executable
	cmd := exec.Command("./" + executablePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
