# Go CyberSec Practice 🛡️ 

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Security Education](https://img.shields.io/badge/Purpose-Educational-green.svg)](#)

This repository contains a collection of mini cybersecurity tools written in Go. The projects are designed to explore the Go programming language alongside core ethical hacking and defensive security concepts, covering cryptography, concurrent networking, system interactions, file integrity, and threat intelligence.

## 🛠️ Projects Included

* **Password Analyzer:** A utility to evaluate password strength. It analyzes length and character complexity using the `unicode` package.
* **MD5 PIN Bruteforcer:** An offline hash cracker. Demonstrates brute-forcing concepts by hashing and iterating through 10,000 PIN combinations to find a match for a target MD5 hash.
* **Concurrent TCP Port Scanner:** An ultra-fast, multi-threaded port scanner using **Goroutines** and `sync.WaitGroup` to scan hundreds of ports concurrently.
* **Web Directory Scanner (Dirbuster clone):** A tool to find hidden directories and files performing HTTP GET requests and analyzing status codes.
* **Interactive C2 Shell (Backdoor Prototype):** A remote command interface that reads commands, parses them, and executes system commands directly using `os/exec`.
* **XOR Payload Obfuscator:** A data encryption tool implementing bitwise XOR encryption to hide sensitive strings from static antivirus analysis.
* **Recursive File Encryptor (Ransomware Simulator):** Explores file system traversal using the `path/filepath` package to read, XOR-encrypt, and overwrite files.
* **Continuous FIM (File Integrity Monitor) Daemon:** A defensive background service. It continuously calculates and monitors the SHA-256 cryptographic hash of critical files using `crypto/sha256` in an infinite loop, alerting administrators immediately upon unauthorized modifications.
* **VirusTotal API Integrator (Threat Intelligence):** A tool that interacts with the VirusTotal v3 API to analyze IP addresses. It securely passes API keys via HTTP headers, handles rate-limiting and authentication HTTP status codes, and unmarshals complex JSON responses to extract security engine detection stats.

## 🚀 How to Run

Ensure you have [Go](https://go.dev/dl/) installed on your machine.
Navigate to the directory of the specific project you want to test and execute:

```bash
go run <filename>.go
