package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "io"
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "strings"
)

var key = []byte("passphrasewhichneedstobe32bytes!") // 32 bytes key for AES-256

func encrypt(data []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    
    return gcm.Seal(nonce, nonce, data, nil), nil
}

func decrypt(data []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonceSize := gcm.NonceSize()
    if len(data) < nonceSize {
        return nil, err
    }
    
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    return gcm.Open(nil, nonce, ciphertext, nil)
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Usage: ./main [-e] <python_script_path>")
    }

    isEncrypt := os.Args[1] == "-e"
    scriptPath := os.Args[1]
    if isEncrypt {
        if len(os.Args) < 3 {
            log.Fatal("No script provided for encryption")
        }
        scriptPath = os.Args[2]
    }

    // Check if file exists
    content, err := ioutil.ReadFile(scriptPath)
    if err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    if isEncrypt {
        // Encrypt mode
        encrypted, err := encrypt(content)
        if err != nil {
            log.Fatalf("Error encrypting: %v", err)
        }

        outPath := strings.TrimSuffix(scriptPath, ".py") + "_encrypted.py"
        err = ioutil.WriteFile(outPath, []byte(hex.EncodeToString(encrypted)), 0644)
        if err != nil {
            log.Fatalf("Error writing encrypted file: %v", err)
        }
        log.Printf("Encrypted file saved as: %s", outPath)
        return
    }

    // Decrypt mode
    encrypted, err := hex.DecodeString(string(content))
    if err != nil {
        log.Fatalf("Error decoding: %v", err)
    }

    decrypted, err := decrypt(encrypted)
    if err != nil {
        log.Fatalf("Error decrypting: %v", err)
    }

    // Create temporary file for decrypted script
    tmpFile, err := ioutil.TempFile("", "script_*.py")
    if err != nil {
        log.Fatalf("Error creating temp file: %v", err)
    }
    defer os.Remove(tmpFile.Name())

    if _, err := tmpFile.Write(decrypted); err != nil {
        log.Fatalf("Error writing temp file: %v", err)
    }
    tmpFile.Close()

    // Execute decrypted script
    pythonCmd := "python3"
    if _, err := exec.LookPath(pythonCmd); err != nil {
        pythonCmd = "python"
        if _, err := exec.LookPath(pythonCmd); err != nil {
            log.Fatal("Python interpreter not found")
        }
    }

    cmd := exec.Command(pythonCmd, tmpFile.Name())
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        log.Fatalf("Error running Python script: %v", err)
    }
}