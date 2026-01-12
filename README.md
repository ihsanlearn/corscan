<h1 align="center">
  <br>
  <a href="https://github.com/iihsannlearn/corscan"><img src="https://img.shields.io/badge/Corscan-CORS%20Vulnerability%20Scanner-blue?style=for-the-badge&logo=go" alt="Corscan"></a>
  <br>
</h1>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#advanced-usage-waf-evasion">Advanced Usage</a> •
  <a href="#license">License</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/language-Go-blue?style=flat-square" alt="Language">
  <img src="https://img.shields.io/badge/license-MIT-green?style=flat-square" alt="License">
  <img src="https://img.shields.io/github/issues/iihsannlearn/corscan?style=flat-square" alt="Issues">
</p>

<p align="center">
  <b>Corscan</b> is a professional, high-performance CORS (Cross-Origin Resource Sharing) vulnerability scanner written in Go. It is designed for speed, flexibility, and stealth, making it an essential tool for security researchers and bug hunters.
</p>

## Features

- **Fast & Multi-threaded**: Scans hundreds of URLs concurrently.
- **WAF Evasion**: Built-in techniques to bypass Web Application Firewalls (Random User-Agents, Origin Shifting, Proxy support).
- **Stealth Mode**: Automated header manipulation to stay undetected.
- **Smart Payloads**: Includes default CORS misconfiguration payloads and supports custom payload lists.
- **Advanced Logic**: Detects reflected Origins, `null` origins, and credentials support (`Access-Control-Allow-Credentials`).
- **Input Flexibility**: Accepts URLs via flag, file, or stdin (pipe).
- **Professional Logger**: Clear, color-coded output.

## Installation

Ensure you have **Go 1.20+** installed.

```bash
go install github.com/iihsannlearn/corscan@latest
```

Or build from source:

```bash
git clone https://github.com/iihsannlearn/corscan
cd corscan
go build -o corscan ./cmd/corscan
```

## Usage

### Basic Scan

Scan a single URL:

```bash
corscan -u http://example.com
```

Scan a list of URLs from a file:

```bash
corscan -l urls.txt
```

Scan from stdin (useful for chaining commands):

```bash
cat urls.txt | corscan
```

### Full Options

```console
Flags:
INPUT:
   -u, -url string      Target URL for scanning (comma separated)
   -l, -list string     File containing list of target URLs
   -or, -origin string  Custom origin for testing

CONFIGURATION:
   -t, -threads int      Number of concurrent threads (default 25)
   -T, -timeout int      Timeout request (seconds) (default 10)
   -rl, -rate-limit int  Maximum requests per second (default 10)
   -d, -delay int        Delay between requests (milliseconds)
   -x, -proxy string     HTTP Proxy URL (e.g. http://127.0.0.1:8080)

OUTPUT:
   -o, -output string  File for saving scan results

OPTIMIZATION:
   -ra, -random-agent  Enable Random User-Agent
   -v, -verbose        Verbose mode (print error & debug)
   -V, -version        Display application version
```

## Advanced Usage (WAF Evasion)

**Corscan** comes with powerful flags to evade detection and bypass filters.

### 1. Custom Origin

Test a specific origin (e.g., subdomain or internal domain) to check for trust relationships.

```bash
corscan -u http://target.com -or internal.target.com
```

### 2. User-Agent Rotation

Use `-random-agent` (`-ra`) to rotate User-Agents for every request, preventing blocking based on static signatures.

```bash
corscan -l targets.txt -ra
```

### 3. Proxy Support

Route traffic through a proxy (e.g., Burp Suite, Tor) for analysis or anonymity.

```bash
corscan -u http://target.com -x http://127.0.0.1:8080
```

### 4. Stealth Tactics (Automatic)

When scanning, **Corscan** automatically attempts:

- **Origin Shifting**: Tries variations like `null`, `http://evil.com`, and custom origins.
- **HTTP Parameter Pollution**: Injects multiple headers to confuse WAFs.
- **Reverse Headers**: Injects spoofed `X-Forwarded-For` and `X-Real-IP` headers with random IPs.

---
