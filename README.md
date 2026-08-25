# Go CyberSec Practice 🛡️ 

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Security Education](https://img.shields.io/badge/Purpose-Educational-green.svg)](#)

This repository contains a collection of mini cybersecurity tools written in Go. The projects are designed to explore the Go programming language alongside core ethical hacking concepts, covering cryptography, concurrent networking, and system interactions.

## 🛠️ Projects Included

* **Password Analyzer:** A utility to evaluate password strength. It analyzes length and character complexity (digits, special characters, uppercase) using the `unicode` package.
* **MD5 PIN Bruteforcer:** An offline hash cracker. Demonstrates brute-forcing concepts by hashing and iterating through 10,000 PIN combinations (0000 to 9999) to find a match for a target MD5 hash using the `crypto/md5` package.
* **Concurrent TCP Port Scanner:** An ultra-fast, multi-threaded port scanner. Utilizes **Goroutines** and `sync.WaitGroup` to scan hundreds of ports concurrently in seconds, effectively bypassing standard network timeout delays.
* **Web Directory Scanner (Dirbuster clone):** A tool to find hidden directories and files (e.g., `/admin`, `/backup.zip`). It performs HTTP GET requests and analyzes status codes (200 OK) to discover non-public paths on a web server.
* **Interactive C2 Shell (Backdoor Prototype):** A remote command interface. It reads commands with arguments via `bufio`, parses them, and executes system commands directly using `os/exec`, returning the output to the terminal.
* **XOR Payload Obfuscator:** A data encryption tool. Implements bitwise XOR encryption to hide sensitive strings (like C2 URLs or payloads) from static antivirus analysis, demonstrating in-memory payload decryption.
* **Recursive File Encryptor (Ransomware Simulator):** Explores file system traversal using the `path/filepath` package. It recursively iterates through a specified directory, reading, XOR-encrypting, and overwriting files to demonstrate how ransomware operates at a fundamental level.

## 🚀 How to Run

Ensure you have [Go](https://go.dev/dl/) installed on your machine.
Navigate to the directory of the specific project you want to test and execute:

```bash
go run <filename>.go
