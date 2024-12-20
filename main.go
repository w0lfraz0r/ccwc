package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
    // Define command-line flags
    countBytes := flag.Bool("c", false, "Count bytes in the file")
    countLines := flag.Bool("l", false, "Count newlines in the file")
    countWords := flag.Bool("w", false, "Count words in the file")

	flag.Parse()

    // Check if a file is provided
    if len(flag.Args()) == 0 {
        fmt.Println("Usage: ccwc [OPTIONS] [FILE]")
        os.Exit(1)
    }

    // Read the file
    filename := flag.Args()[0]
    content, err := os.ReadFile(filename)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }

    // Variables to hold counts
    var byteCount, lineCount, wordCount int
    output := ""

    // Handle -c flag
    if *countBytes {
        byteCount = len(content)
        output += fmt.Sprintf("%8d ", byteCount)
    }

    // Handle -l flag
    if *countLines {
        file, err := os.Open(filename)
        if err != nil {
            fmt.Printf("Error opening file: %v\n", err)
            os.Exit(1)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            lineCount++
        }

        if err := scanner.Err(); err != nil {
            fmt.Printf("Error reading file: %v\n", err)
            os.Exit(1)
        }

        output += fmt.Sprintf("%8d ", lineCount)
    }

	// Handle -w flag
if *countWords {
        words := strings.Fields(string(content))
        wordCount = len(words)
        output += fmt.Sprintf("%8d ", wordCount)
    }

    output += filename
    // Printing Results
    if output != filename {
		fmt.Println(output)
	}
}
