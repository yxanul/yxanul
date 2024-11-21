package main

import (
	"os"
	"os/exec"
	"log"
)

func RunHTTPX(inputFile string) {
	httpxArgs := []string{"-l", inputFile, "-silent", "-nc", "-o", "httpx_output.txt"}
	cmd := exec.Command("httpx", httpxArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running HTTPX: %v", err)
	}
}

