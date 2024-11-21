package main

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

// RunWAFDetection checks for WAF presence using wafw00f and extracts the detection result.
func RunWAFDetection(url string) (bool, string, error) {
	// Run the wafw00f command
	cmd := exec.Command("wafw00f", url)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	if err := cmd.Run(); err != nil {
		return false, "", err
	}

	// Parse the wafw00f output
	wafDetected := false
	detectionMessage := "No WAF detected"
	scanner := bufio.NewScanner(&output)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Check for specific output indicating no WAF
		if strings.Contains(line, "No WAF detected by the generic detection") {
			wafDetected = false
			detectionMessage = "No WAF detected"
			break
		}

		// Check for specific WAF detection messages
		if strings.Contains(line, "is behind") {
			wafDetected = true
			detectionMessage = line // Example: "The site https://example.com is behind Cloudflare (Cloudflare Inc.) WAF."
			break
		} else if strings.Contains(line, "WAF detected") {
			wafDetected = true
			detectionMessage = "Generic WAF detected" // Default fallback
			break
		} else if strings.Contains(line, "seems to be behind a WAF") {
			wafDetected = true
			detectionMessage = line // Example: "The site seems to be behind a WAF or security solution."
			break
		}
	}

	// If no WAF-specific message is found and the default detection is unclear, assume no WAF
	return wafDetected, detectionMessage, nil
}
