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