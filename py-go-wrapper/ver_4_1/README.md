# Plan to Update Deprecated io/ioutil
1. Replace with `os` package functions
2. Update file operations to use new APIs
3. Remove deprecated import

```go


package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "io"
    "log"
    "os"
    "os/exec"
    "strings"
    "bytes"
)

// ...existing encryption/decryption functions...

func main() {
    // ...existing argument checks...

    content, err := os.ReadFile(scriptPath)
    if err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    if isEncrypt {
        encrypted, err := encrypt(content)
        if err != nil {
            log.Fatalf("Error encrypting: %v", err)
        }

        outPath := strings.TrimSuffix(scriptPath, ".py") + "_encrypted.py"
        err = os.WriteFile(outPath, []byte(hex.EncodeToString(encrypted)), 0644)
        if err != nil {
            log.Fatalf("Error writing encrypted file: %v", err)
        }
        log.Printf("Encrypted file saved as: %s", outPath)
        return
    }

    // ...existing encryption check...

    tmpFile, err := os.CreateTemp("", "script_*.py")
    if err != nil {
        log.Fatalf("Error creating temp file: %v", err)
    }
    defer os.Remove(tmpFile.Name())

    if _, err := tmpFile.Write(decrypted); err != nil {
        log.Fatalf("Error writing temp file: %v", err)
    }
    tmpFile.Close()

    // ...existing execution code...
}
```