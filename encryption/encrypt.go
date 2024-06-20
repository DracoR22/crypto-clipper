package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func EncryptFile(filePath string, encryptionKey []byte) ([]byte, error) {
	// Open the file to encrypt
	inputFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	// Create a temporary file to store encrypted data
	encryptedFilePath := "main.enc.tmp"
	outputFile, err := os.Create(encryptedFilePath)
	if err != nil {
		return nil, err
	}
	defer outputFile.Close()

	// Generate a random IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Write IV to the beginning of the output file
	if _, err := outputFile.Write(iv); err != nil {
		return nil, err
	}

	// Create AES cipher block
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	// Create CBC mode encrypter
	mode := cipher.NewCBCEncrypter(block, iv)

	// Buffer for reading input file
	buffer := make([]byte, 4096)
	for {
		// Read a chunk of data from input file
		bytesRead, err := inputFile.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if bytesRead == 0 {
			break
		}

		// Pad the last block if needed
		plaintext := buffer[:bytesRead]
		paddedPlaintext := padData(plaintext, aes.BlockSize)

		// Encrypt the chunk
		encryptedChunk := make([]byte, len(paddedPlaintext))
		mode.CryptBlocks(encryptedChunk, paddedPlaintext)

		// Write the encrypted chunk to the output file
		if _, err := outputFile.Write(encryptedChunk); err != nil {
			return nil, err
		}
	}

	// Close the output file before renaming
	if err := outputFile.Close(); err != nil {
		return nil, err
	}

	// Rename the temporary file to the final encrypted file name
	if err := os.Rename(encryptedFilePath, "main.exe.enc"); err != nil {
		return nil, err
	}

	return iv, nil
}

func padData(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
