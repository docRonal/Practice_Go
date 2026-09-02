# Go CyberSec Practice 🛡️ 

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Security Education](https://img.shields.io/badge/Purpose-Educational-green.svg)](#)

# ⚠️ Disclaimer

**EDUCATIONAL PURPOSES ONLY.**

**All tools and scripts in this repository are provided strictly for educational purposes and ethical hacking practice. They are designed to help understand system vulnerabilities and
  improve defensive security postures.**

  * Do NOT use these tools on any system, network, or application that you do not own or do not have explicit, documented permission to test.

  * The author assumes no responsibility and shall not be held liable for any illegal, malicious, or unauthorized use of this code or any damage caused by it.


This repository contains a collection of mini cybersecurity tools written in Go. The projects are designed to explore the Go programming language alongside core ethical hacking and defensive security concepts, covering cryptography, concurrent networking, system interactions, file integrity, and threat intelligence API integrations.

## 🛠️ Projects Included

* **Password Analyzer:** A utility to evaluate password strength. It analyzes length and character complexity using the `unicode` package.
* **MD5 PIN Bruteforcer:** An offline hash cracker. Demonstrates brute-forcing concepts by hashing and iterating through 10,000 PIN combinations to find a match for a target MD5 hash.
* **Concurrent TCP Port Scanner:** An ultra-fast, multi-threaded port scanner using **Goroutines** and `sync.WaitGroup` to scan hundreds of ports concurrently.
* **Web Directory Scanner (Dirbuster clone):** A tool to find hidden directories and files performing HTTP GET requests and analyzing status codes.
* **Interactive C2 Shell (Backdoor Prototype):** A remote command interface that reads commands, parses them, and executes system commands directly using `os/exec`.
* **XOR Payload Obfuscator:** A data encryption tool implementing bitwise XOR encryption to hide sensitive strings from static antivirus analysis.
* **Recursive File Encryptor (Ransomware Simulator):** Explores file system traversal using the `path/filepath` package to read, XOR-encrypt, and overwrite files.
* **Continuous FIM (File Integrity Monitor) Daemon:** A defensive background service. It continuously calculates and monitors the SHA-256 cryptographic hash of critical files using `crypto/sha256` in an infinite loop, alerting administrators immediately upon unauthorized modifications.
* **Threat Intelligence IP Analyzer:** An OSINT tool that queries the VirusTotal API (`v3/ip_addresses`) to retrieve reputation data. It parses nested JSON responses using `encoding/json` and implements robust HTTP status code checking for rate-limiting and authentication errors.
* **Malware Hash Analyzer:** A Threat Intelligence utility that checks SHA-256 file hashes against the VirusTotal API (`v3/files`) to determine if a file is flagged as malicious by leading antivirus engines.
* **Web Server Log Analyzer (IDS/WAF Prototype):** A Blue Team defensive tool that parses HTTP web server logs (e.g., `access.log`). It uses the `regexp` package to detect common web application attacks such as SQL Injections, Cross-Site Scripting (XSS), and Path Traversal attempts in real-time.
## 🚀 How to Run

Ensure you have [Go](https://go.dev/dl/) installed on your machine.
Navigate to the directory of the specific project you want to test and execute:

```bash
go run <filename>.go
