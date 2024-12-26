# Plan for Running Python Script with Go

1. Create Go wrapper that:
   - Executes Python interpreter
   - Passes arguments/parameters 
   - Handles output/errors
   - Manages Python process

2. Implementation steps:
   - Set up Python environment detection
   - Create Go executable
   - Handle Python script execution
   - Manage return values

```go
package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
)

func main() {
    // Python script to run
    pythonScript := "script.py"

    // Check if Python exists
    pythonCmd := "python3"
    if _, err := exec.LookPath(pythonCmd); err != nil {
        pythonCmd = "python"
        if _, err := exec.LookPath(pythonCmd); err != nil {
            log.Fatal("Python interpreter not found")
        }
    }

    // Create command
    cmd := exec.Command(pythonCmd, pythonScript)
    
    // Set up pipes for output
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Run the command
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Error running Python script: %v", err)
    }
}
```

```python
def main():
    print("Hello from Python!")
    return 0

if __name__ == "__main__":
    main()
```

3. Build and run:
```bash
go build main.go
./main
```

This creates a Go executable that runs the Python script. The Go program:
- Detects Python installation
- Executes the script
- Handles output/errors
- Returns exit code


# Plan: Modify Go Program to Accept Python Script as CLI Argument

1. Add argument checking
2. Use os.Args to get script path
3. Add usage message
4. Handle missing arguments

```go


package main

import (
    "log"
    "os"
    "os/exec"
)

func main() {
    // Check if script path is provided
    if len(os.Args) < 2 {
        log.Fatal("Usage: ./main <python_script_path>")
    }

    // Get script path from command line argument
    pythonScript := os.Args[1]

    // Check if file exists
    if _, err := os.Stat(pythonScript); os.IsNotExist(err) {
        log.Fatalf("Python script not found: %s", pythonScript)
    }

    pythonCmd := "python3"
    if _, err := exec.LookPath(pythonCmd); err != nil {
        pythonCmd = "python"
        if _, err := exec.LookPath(pythonCmd); err != nil {
            log.Fatal("Python interpreter not found")
        }
    }

    cmd := exec.Command(pythonCmd, pythonScript)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        log.Fatalf("Error running Python script: %v", err)
    }
}
```

Usage:
```bash
go build main.go
./main script.py
```


# Plan for Python Script Encryption/Decryption Wrapper

1. Steps:
   - Add flag parsing for encryption mode
   - Create encryption/decryption functions using AES
   - Handle file operations (read/write)
   - Process Python script based on mode
   - Execute decrypted script
   - Clean up temporary files

2. Implementation:

```go


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
    "path/filepath"
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
```

3. Usage:
```bash
# To encrypt a Python script
go build main.go
./main -e script.py

# To run an encrypted script
./main script_encrypted.py
```

Similar code found with 3 license types

# Plan for Fixing Hex Decoding Error

1. Check if file is encrypted
2. Add validation before decoding
3. Execute unencrypted files directly
4. Better error handling

```go
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
    "bytes"
)

// ... existing encryption/decryption functions ...

func isEncryptedFile(content []byte) bool {
    // Check if content is hex encoded
    _, err := hex.DecodeString(string(content))
    return err == nil && !bytes.Contains(content, []byte{' '})
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

    content, err := ioutil.ReadFile(scriptPath)
    if err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    if isEncrypt {
        // Encryption mode
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

    // Check if file is encrypted
    if !isEncryptedFile(content) {
        // Run unencrypted file directly
        pythonCmd := "python3"
        if _, err := exec.LookPath(pythonCmd); err != nil {
            pythonCmd = "python"
            if _, err := exec.LookPath(pythonCmd); err != nil {
                log.Fatal("Python interpreter not found")
            }
        }

        cmd := exec.Command(pythonCmd, scriptPath)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error running Python script: %v", err)
        }
        return
    }

    // Decrypt and run encrypted file
    encrypted, err := hex.DecodeString(string(content))
    if err != nil {
        log.Fatalf("Error decoding encrypted file: %v", err)
    }

    decrypted, err := decrypt(encrypted)
    if err != nil {
        log.Fatalf("Error decrypting: %v", err)
    }

    tmpFile, err := ioutil.TempFile("", "script_*.py")
    if err != nil {
        log.Fatalf("Error creating temp file: %v", err)
    }
    defer os.Remove(tmpFile.Name())

    if _, err := tmpFile.Write(decrypted); err != nil {
        log.Fatalf("Error writing temp file: %v", err)
    }
    tmpFile.Close()

    cmd := exec.Command("python3", tmpFile.Name())
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        log.Fatalf("Error running decrypted script: %v", err)
    }
}
```