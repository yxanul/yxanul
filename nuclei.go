package main

import (
	"log"
	"os"
	"os/exec"
)

// RunNuclei runs Nuclei on the given input file (URLs) and applies appropriate options.
func RunNuclei(inputFile string) {
	// Base arguments for Nuclei
	nucleiArgs := []string{"-l", inputFile, "-silent", "-as", "-stats", "-o", "nuclei_output.txt"}

	// Execute Nuclei command
	cmd := exec.Command("nuclei", nucleiArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Handle errors during execution
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running Nuclei: %v", err)
	}
}
