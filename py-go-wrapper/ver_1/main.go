package main

import (
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