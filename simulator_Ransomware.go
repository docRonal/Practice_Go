package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Use a safe test folder so we don't accidentally modify source files.
	startDir := "./test_files"
	key := "hacker"
	keyBytes := []byte(key)

	err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		// If Walk encountered an error with this path, skip it to avoid nil deref on info
		if err != nil {
			fmt.Println("walk error:", err)
			return nil
		}
		if info == nil {
			return nil
		}

		if !info.IsDir() {
			data, readErr := os.ReadFile(path)
			if readErr != nil {
				fmt.Println("Error reading file:", readErr)
				return nil
			}

			for i := 0; i < len(data); i++ {
				data[i] = data[i] ^ keyBytes[i%len(keyBytes)]
			}

			writeErr := os.WriteFile(path, data, 0644)
			if writeErr != nil {
				fmt.Println("Error writing file:", writeErr)
				return writeErr
			}

			fmt.Println("file encoded", path)
		} else {
			fmt.Println("Ignoring directory:", path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error", err)
	}
}
