package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Define command-line flags
	url := flag.String("url", "", "Target URL to scan")
	aggressive := flag.Bool("a", false, "Run in aggressive mode")
	flag.Parse()

	// Validate input
	if *url == "" {
		fmt.Println("Error: Please specify a URL using the -url flag")
		return
	}

	fmt.Println(`


   Web Scanning Toolkit by <yxanul>
-----------------------------------------------------
	`)

	// Step 1: Run WAF detection
	// Run WAF detection
fmt.Println("Checking for WAF presence...")
wafDetected, wafInfo, err := RunWAFDetection(*url)
if err != nil {
	log.Fatalf("Error running WAF detection: %v", err)
}

// Print WAF detection results
if wafDetected {
	fmt.Printf("WAF Detected: %s\n", wafInfo)
} else {
	fmt.Println("No WAF detected.")
}


	// Step 2: Run Katana with appropriate configuration
	fmt.Println("Running Katana for crawling...")
	RunKatana(*url, wafDetected, *aggressive)

	// Step 3: Run HTTPX to probe URLs from Katana's output
	fmt.Println("Running HTTPX to probe URLs...")
	RunHTTPX("katana_output.txt")

	// Step 4: Run Nuclei for vulnerability scanning on probed URLs
	fmt.Println("Running Nuclei on probed URLs...")
	RunNuclei("httpx_output.txt")

	// Final message
	fmt.Println("\nScan completed successfully!")
}
