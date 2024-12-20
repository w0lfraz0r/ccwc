package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
    // Define command-line flags
    countBytes := flag.Bool("c", false, "Count bytes in the file")
    countLines := flag.Bool("l", false, "Count newlines in the file")
    countWords := flag.Bool("w", false, "Count words in the file")
    countChars := flag.Bool("m", false, "Count characters in the file")

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
    byteCount := len(content)
    lineCount := 0
    wordCount := 0
    charCount := utf8.RuneCount(content)

    // Calculate line count
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

    // Calculate word count
    wordCount = len(strings.Fields(string(content)))

    // Detect default output
    defaultOutput := !(*countBytes || *countLines || *countWords || *countChars)

    // Prepare output
    output := ""
    if *countLines || defaultOutput {
        output += fmt.Sprintf("%8d ", lineCount)
    }
    if *countWords || defaultOutput {
        output += fmt.Sprintf("%8d ", wordCount)
    }
    if *countBytes || defaultOutput {
        output += fmt.Sprintf("%8d ", byteCount)
    }
    if *countChars {
        output += fmt.Sprintf("%8d ", charCount)
    }

    output += filename
    fmt.Println(output)
}
