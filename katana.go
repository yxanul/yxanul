package main

import (
	"log"
	"os"
	"os/exec"
)

func RunKatana(url string, wafDetected bool, aggressive bool) {
	// Base arguments for all scenarios
	katanaArgs := []string{"-u", url, "-silent", "-nc", "-o", "katana_output.txt", "-jc", "-jsl", "-hl", "-nos", "-f", "qurl"}

	// Adjust arguments based on WAF detection and aggressiveness
	if wafDetected {
		// WAF detected scenarios
		if aggressive {
			// Aggressive WAF detected mode
			katanaArgs = append(katanaArgs, "-c", "15", "-p", "15", "-rl", "20", "-d", "5")
		} else {
			// Default WAF detected mode
			katanaArgs = append(katanaArgs, "-c", "12", "-p", "12", "-rl", "10", "-d", "3")
		}
	} else {
		// No WAF scenarios
		if aggressive {
			// Aggressive no WAF mode
			katanaArgs = append(katanaArgs, "-c", "30", "-p", "20", "-rl", "500", "-d", "5")
		} else {
			// Default no WAF mode
			katanaArgs = append(katanaArgs, "-c", "15", "-p", "10", "-rl", "50", "-d", "3")
		}
	}

	// Run Katana command
	cmd := exec.Command("katana", katanaArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running Katana: %v", err)
	}
}
