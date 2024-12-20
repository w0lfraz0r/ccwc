package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
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

    // Read the file or standard input
    var content []byte
    var err error
    var filename string
    if len(flag.Args()) == 0 {
        content, err = ioutil.ReadAll(os.Stdin)
        if err != nil {
            fmt.Printf("Error reading standard input: %v\n", err)
            os.Exit(1)
        }
    } else {
        filename = flag.Args()[0]
        content, err = os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
            os.Exit(1)
        }
    }

    // Variables to hold counts
    byteCount := len(content)
    lineCount := 0
    wordCount := 0
    charCount := utf8.RuneCount(content)

    // Calculate line count
    if filename != "" {
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
    } else {
        lineCount = strings.Count(string(content), "\n")
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

    // Print the results
    if filename != "" {
        output += filename
    }
    fmt.Println(output)
}
