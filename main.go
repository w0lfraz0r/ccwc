package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "os"
)

func main() {
    // Define command-line flags
    countBytes := flag.Bool("c", false, "Count bytes in the file")

    flag.Parse()

    // Check if a file is provided
    if len(flag.Args()) == 0 {
        fmt.Println("Usage: ccwc [OPTIONS] [FILE]")
        os.Exit(1)
    }

    // Read the file
    filename := flag.Args()[0]
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }

    // Handle -c flag
    if *countBytes {
        fmt.Printf("%8d %s\n", len(content), filename)
    }
}
