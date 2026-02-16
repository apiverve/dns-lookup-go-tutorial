// DNS Lookup Tool - Tutorial Example
//
// A simple CLI tool that looks up DNS records using the APIVerve API.
// https://apiverve.com/marketplace/dnslookup
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// ============================================
// CONFIGURATION - Add your API key here
// Get a free key at: https://dashboard.apiverve.com
// ============================================
const (
	apiKey = "your-api-key-here"
	apiURL = "https://api.apiverve.com/v1/dnslookup"
)

// APIResponse represents the API response structure
type APIResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   APIData     `json:"data"`
}

// APIData contains the domain and records
type APIData struct {
	Domain  string   `json:"domain"`
	Records DNSData  `json:"records"`
}

// DNSData contains the DNS record information
type DNSData struct {
	A     []string   `json:"A"`
	AAAA  []string   `json:"AAAA"`
	MX    []MXRecord `json:"MX"`
	TXT   []string   `json:"TXT"`
	CNAME []string   `json:"CNAME"`
	NS    []string   `json:"NS"`
}

// MXRecord represents a mail exchange record
type MXRecord struct {
	Exchange string `json:"exchange"`
	Priority int    `json:"priority"`
}

// lookupDNS queries the DNS records for a domain
func lookupDNS(domain string) (*DNSData, error) {
	if apiKey == "your-api-key-here" {
		return nil, fmt.Errorf("API key not configured. Add your key to main.go")
	}

	// Clean domain
	domain = strings.TrimSpace(strings.ToLower(domain))
	domain = strings.TrimPrefix(domain, "https://")
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.Split(domain, "/")[0]

	// Create request
	req, err := http.NewRequest("GET", apiURL+"?domain="+domain, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("x-api-key", apiKey)

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Parse response
	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if apiResp.Status != "ok" {
		errMsg := apiResp.Error
		if errMsg == "" {
			errMsg = "DNS lookup failed"
		}
		return nil, fmt.Errorf(errMsg)
	}

	return &apiResp.Data.Records, nil
}

// printRecords displays DNS records in a formatted way
func printRecords(domain string, data *DNSData) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 55))
	fmt.Printf("  DNS Records: %s\n", domain)
	fmt.Println(strings.Repeat("=", 55))

	// A Records (IPv4)
	fmt.Println("\n  A Records (IPv4)")
	fmt.Println("  " + strings.Repeat("-", 51))
	if len(data.A) > 0 {
		for _, record := range data.A {
			fmt.Printf("    %s\n", record)
		}
	} else {
		fmt.Println("    (none)")
	}

	// AAAA Records (IPv6)
	fmt.Println("\n  AAAA Records (IPv6)")
	fmt.Println("  " + strings.Repeat("-", 51))
	if len(data.AAAA) > 0 {
		for _, record := range data.AAAA {
			fmt.Printf("    %s\n", record)
		}
	} else {
		fmt.Println("    (none)")
	}

	// MX Records (Mail)
	fmt.Println("\n  MX Records (Mail)")
	fmt.Println("  " + strings.Repeat("-", 51))
	if len(data.MX) > 0 {
		for _, record := range data.MX {
			fmt.Printf("    [%d] %s\n", record.Priority, record.Exchange)
		}
	} else {
		fmt.Println("    (none)")
	}

	// NS Records (Nameservers)
	fmt.Println("\n  NS Records (Nameservers)")
	fmt.Println("  " + strings.Repeat("-", 51))
	if len(data.NS) > 0 {
		for _, record := range data.NS {
			fmt.Printf("    %s\n", record)
		}
	} else {
		fmt.Println("    (none)")
	}

	// TXT Records
	fmt.Println("\n  TXT Records")
	fmt.Println("  " + strings.Repeat("-", 51))
	if len(data.TXT) > 0 {
		for _, record := range data.TXT {
			// Truncate long TXT records
			if len(record) > 60 {
				record = record[:60] + "..."
			}
			fmt.Printf("    %s\n", record)
		}
	} else {
		fmt.Println("    (none)")
	}

	// CNAME Records
	if len(data.CNAME) > 0 {
		fmt.Println("\n  CNAME Records")
		fmt.Println("  " + strings.Repeat("-", 51))
		for _, record := range data.CNAME {
			fmt.Printf("    %s\n", record)
		}
	}

	fmt.Println()
	fmt.Println(strings.Repeat("=", 55))
	fmt.Println()
}

// interactiveMode runs the tool in interactive mode
func interactiveMode() {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 55))
	fmt.Println("  DNS Lookup Tool")
	fmt.Println("  Powered by APIVerve")
	fmt.Println(strings.Repeat("=", 55))
	fmt.Println("\nLook up DNS records for any domain")
	fmt.Println("Type 'quit' to exit\n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter domain (e.g., google.com): ")
		if !scanner.Scan() {
			break
		}

		domain := strings.TrimSpace(scanner.Text())
		if strings.ToLower(domain) == "quit" {
			break
		}

		if domain == "" {
			fmt.Println("Please enter a domain.\n")
			continue
		}

		data, err := lookupDNS(domain)
		if err != nil {
			fmt.Printf("\n  ❌ Error: %v\n\n", err)
			continue
		}

		printRecords(domain, data)
	}

	fmt.Println("Goodbye!")
}

func main() {
	if len(os.Args) == 2 {
		// Command line mode: go run main.go google.com
		domain := os.Args[1]
		data, err := lookupDNS(domain)
		if err != nil {
			fmt.Printf("\n  ❌ Error: %v\n\n", err)
			os.Exit(1)
		}
		printRecords(domain, data)
	} else {
		// Interactive mode
		interactiveMode()
	}
}
