# ccwc - Custom Word Count Tool

`ccwc` is a custom implementation of the Unix `wc` command, built in Go. It provides word, line, character, and byte counts for text files or input streams.

## Features

- Count the number of **bytes** in a file using `-c`.
- Count the number of **lines** in a file using `-l`.
- Count the number of **words** in a file using `-w`.
- Count the number of **characters** in a file using `-m`.
- Supports multiple options and combines outputs.
- Works with standard input when no filename is provided.

## Installation

### Prerequisites

- Go 1.20 or later installed on your system.
- A Unix-based operating system (developed on Fedora).

### Clone and Build

1. Clone the repository:
   ```bash
   git clone https://github.com/w0lfraz0r/ccwc.git
   cd ccwc
   ```
2. Build the project:
   ```bash
   go build -o ccwc
   ```
3. Run the tool:
   ```bash
   ./ccwc -c test.txt
   ```

## Usage

Basic syntax:

```bash
ccwc [OPTIONS] [FILE]
```

### Options

- `-c` : Count the number of bytes.
- `-l` : Count the number of lines.
- `-w` : Count the number of words.
- `-m` : Count the number of characters.
- _(default)_ : Outputs all counts (bytes, lines, words).

### Examples

1. Count lines in a file:
   ```bash
   ccwc -l test.txt
   ```
2. Count words and characters:
   ```bash
   ccwc -w -m test.txt
   ```
3. Use standard input:
   ```bash
   cat test.txt | ccwc -l
   ```

## License

This project is licensed under the MIT License. See `LICENSE` for details.

## Acknowledgments

Inspired by the Unix `wc` command and the Unix Philosophy.
