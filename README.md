# DNS Lookup Tool | APIVerve API Tutorial

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Build](https://img.shields.io/badge/Build-Passing-brightgreen.svg)]()
[![Go](https://img.shields.io/badge/Go-1.21+-00add8)](https://go.dev)
[![APIVerve | DNS Lookup](https://img.shields.io/badge/APIVerve-DNS_Lookup-purple)](https://apiverve.com/marketplace/dnslookup?utm_source=github&utm_medium=tutorial&utm_campaign=dns-lookup-go-tutorial)

A Go CLI tool that looks up DNS records for any domain. View A, AAAA, MX, TXT, NS, and CNAME records.

![Screenshot](https://raw.githubusercontent.com/apiverve/dns-lookup-go-tutorial/main/screenshot.jpg)

---

### Get Your Free API Key

This tutorial requires an APIVerve API key. **[Sign up free](https://dashboard.apiverve.com?utm_source=github&utm_medium=tutorial&utm_campaign=dns-lookup-go-tutorial)** - no credit card required.

---

## Features

- Look up all DNS record types
- View A records (IPv4 addresses)
- View AAAA records (IPv6 addresses)
- View MX records with priorities
- View NS records (nameservers)
- View TXT records (SPF, DKIM, etc.)
- View CNAME records
- Interactive mode or command-line arguments
- No external dependencies (standard library only)

## Quick Start

1. **Clone this repository**
   ```bash
   git clone https://github.com/apiverve/dns-lookup-go-tutorial.git
   cd dns-lookup-go-tutorial
   ```

2. **Add your API key**

   Open `main.go` and replace the API key:
   ```go
   const apiKey = "your-api-key-here"
   ```

3. **Run the tool**

   Interactive mode:
   ```bash
   go run main.go
   ```

   Command line mode:
   ```bash
   go run main.go google.com
   ```

4. **Build executable (optional)**
   ```bash
   go build -o dns-lookup main.go
   ./dns-lookup google.com
   ```

## Usage Examples

### Look up a domain
```bash
$ go run main.go github.com

=======================================================
  DNS Records: github.com
=======================================================

  A Records (IPv4)
  ---------------------------------------------------
    140.82.112.3
    140.82.112.4

  AAAA Records (IPv6)
  ---------------------------------------------------
    (none)

  MX Records (Mail)
  ---------------------------------------------------
    [1] aspmx.l.google.com
    [5] alt1.aspmx.l.google.com
    [10] alt2.aspmx.l.google.com

  NS Records (Nameservers)
  ---------------------------------------------------
    dns1.p08.nsone.net
    dns2.p08.nsone.net
    dns3.p08.nsone.net
    dns4.p08.nsone.net

  TXT Records
  ---------------------------------------------------
    v=spf1 include:_spf.google.com include:servers...
    MS=ms12345678

=======================================================
```

## Project Structure

```
dns-lookup-go-tutorial/
├── main.go             # Main Go source file
├── go.mod              # Go module file
├── screenshot.jpg      # Preview image
├── LICENSE             # MIT license
├── .gitignore          # Git ignore rules
└── README.md           # This file
```

## How It Works

1. User provides a domain name
2. Go program cleans the input
3. HTTP request sent to DNS Lookup API
4. JSON response parsed into Go structs
5. Records displayed in formatted output

### The API Call

```go
req, _ := http.NewRequest("GET", apiURL+"?domain="+domain, nil)
req.Header.Set("x-api-key", apiKey)

client := &http.Client{}
resp, _ := client.Do(req)
```

## API Reference

**Endpoint:** `GET https://api.apiverve.com/v1/dnslookup`

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `domain` | string | Yes | Domain to look up (e.g., "google.com") |

**Example Response:**

```json
{
  "status": "ok",
  "error": null,
  "data": {
    "domain": "myspace.com",
    "records": {
      "A": ["34.111.176.156"],
      "MX": [
        {"exchange": "us-smtp-inbound-1.mimecast.com", "priority": 10},
        {"exchange": "us-smtp-inbound-2.mimecast.com", "priority": 10}
      ],
      "NS": [
        "ns-cloud-a1.googledomains.com",
        "ns-cloud-a2.googledomains.com"
      ],
      "TXT": ["v=spf1 mx ip4:63.208.226.34 ..."]
    }
  }
}
```

## DNS Record Types

| Type | Description |
|------|-------------|
| A | IPv4 address |
| AAAA | IPv6 address |
| MX | Mail exchange (email servers) |
| NS | Nameservers |
| TXT | Text records (SPF, DKIM, verification) |
| CNAME | Canonical name (alias) |

## Use Cases

- **Network diagnostics** - Troubleshoot DNS issues
- **Email setup** - Verify MX records
- **Security checks** - Review SPF/DKIM records
- **Domain migration** - Verify DNS propagation
- **DevOps** - Automate DNS verification

## Customization Ideas

- Add JSON output format
- Query specific record types only
- Compare DNS across multiple resolvers
- Build a web interface
- Add DNS propagation checking
- Export results to file

## Related APIs

Explore more APIs at [APIVerve](https://apiverve.com/marketplace?utm_source=github&utm_medium=tutorial&utm_campaign=dns-lookup-go-tutorial):

- [SSL Checker](https://apiverve.com/marketplace/sslchecker?utm_source=github&utm_medium=tutorial&utm_campaign=dns-lookup-go-tutorial) - Check SSL certificates
- [WHOIS Lookup](https://apiverve.com/marketplace/whoislookup?utm_source=github&utm_medium=tutorial&utm_campaign=dns-lookup-go-tutorial) - Domain registration info
- [MX Lookup](https://apiverve.com/marketplace/mxlookup?utm_source=github&utm_medium=tutorial&utm_campaign=dns-lookup-go-tutorial) - Mail server lookup

## License

MIT - see [LICENSE](LICENSE)

## Links

- [Get API Key](https://dashboard.apiverve.com?utm_source=github&utm_medium=tutorial&utm_campaign=dns-lookup-go-tutorial) - Sign up free
- [APIVerve Marketplace](https://apiverve.com/marketplace?utm_source=github&utm_medium=tutorial&utm_campaign=dns-lookup-go-tutorial) - Browse 300+ APIs
- [DNS Lookup API](https://apiverve.com/marketplace/dnslookup?utm_source=github&utm_medium=tutorial&utm_campaign=dns-lookup-go-tutorial) - API details
