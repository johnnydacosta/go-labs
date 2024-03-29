package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()
	// Uncomment process() to process sequentially
	//process()

	processConcurrently()

	elapsed := time.Since(start)
	fmt.Printf("Time process: %f seconds\n", elapsed.Seconds())
}

func processConcurrently() {

	fmt.Println("\n---------------------------------------------")
	fmt.Println("Generating AES key")
	key, err := generateSymmetricKeyForAES()

	if err != nil {
		fmt.Println(err)
		return
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Symmetric Key: %x\n", key)
	fmt.Println("---------------------------------------------")
	fmt.Println()

	sourceDirToEncrypt := "/Users/johnnydacosta/Downloads"
	targetEncryptedDir := "/tmp/encrypted"
	ch := make(chan string)
	var waitToFiles int

	err = filepath.Walk(sourceDirToEncrypt, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		waitToFiles += 1
		go func() {
			encryptFileWithStreaming(key, nonce, path, targetEncryptedDir)
			ch <- path
		}()

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	timeout := time.After(5 * time.Second)
	for i := 0; i < waitToFiles; i++ {
		select {
		case result := <-ch:
			fmt.Printf("File %s encrypted\n", result)
		case <-timeout:
			fmt.Println("Timeout, take too long...")
			return
		}
	}
}

func process() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("Generating AES key")
	key, err := generateSymmetricKeyForAES()

	if err != nil {
		fmt.Println(err)
		return
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Symmetric Key: %x\n", key)
	fmt.Println("---------------------------------------------")
	fmt.Println()

	sourceDirToEncrypt := "/Users/johnnydacosta/Downloads"
	targetEncryptedDir := "/tmp/encrypted"

	err = filepath.Walk(sourceDirToEncrypt, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fmt.Printf("File %s encrypted\n", path)
		return encryptFileWithStreaming(key, nonce, path, targetEncryptedDir)
	})

	if err != nil {
		fmt.Println(err)
	}
}

func encryptFileWithStreaming(key, nonce []byte, path, targetEncryptedDir string) error {
	inputFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(targetEncryptedDir + string(os.PathSeparator) + ".." + string(os.PathSeparator) + "encrypted" + string(os.PathSeparator) + filepath.Base(inputFile.Name()) + ".aes")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Créer un bloc AES-GCM
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error creating AES cipher: %v", err)
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("error creating GCM: %v", err)
	}

	outputWriter := bufio.NewWriter(outputFile)
	buf := make([]byte, 4096) // Taille du bloc de lecture

	for {
		// Lecture du bloc suivant depuis le fichier source
		n, err := inputFile.Read(buf)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading input file: %v", err)
		}
		if n == 0 {
			break
		}

		encryptedBlock := aesgcm.Seal(nil, nonce, buf[:n], nil)

		if _, err := outputFile.Write(encryptedBlock); err != nil {
			return fmt.Errorf("error writing to output file: %v", err)
		}
	}

	if err := outputWriter.Flush(); err != nil {
		return fmt.Errorf("erreur when flushing to output file: %v", err)
	}

	return nil
}

func generateSymmetricKeyForAES() ([]byte, error) {
	// Generate a random symmetric key for HMAC and AES
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, errors.New("error generating symmetric key")
	}
	return key, nil
}

func encrypt(key, nonce []byte, message []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating AES cipher: %v", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("error createing GCM %v", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, message, nil)

	return ciphertext, nil
}

func decrypt(key []byte, message []byte) ([]byte, error) {
	// Decrypt data with AES-GCM
	nonce, ciphertext := message[:12], message[12:]
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating AES cipher for decryption: %v", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("error creating GCM for decryption: %v", err)
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("error decrypting data: %v", err)
	}

	return plaintext, nil
}

func encryptFile(key, nonce []byte, path, targetEncryptedDir string) error {

	inputFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(targetEncryptedDir + string(os.PathSeparator) + ".." + string(os.PathSeparator) + "encrypted" + string(os.PathSeparator) + filepath.Base(inputFile.Name()) + ".aes")

	if err != nil {
		return err
	}

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		cipherLine, err := encrypt(key, nonce, []byte(line))
		if err != nil {
			return fmt.Errorf("error when cipher the line: %v", err)
		}

		_, err = fmt.Fprintln(outputWriter, cipherLine)
		if err != nil {
			return fmt.Errorf("error writing to output file: %v", err)
		}
	}

	if err := outputWriter.Flush(); err != nil {
		return fmt.Errorf("erreur when flushing to output file: %v", err)
	}

	return nil
}
