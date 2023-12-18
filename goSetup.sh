#! /usr/bin/env bash

if [[ -z "$1" ]]; then
    echo "No project name provided"
    exit 1
fi

projName="$1"

mkdir -p "$projName"

cd "$projName"

touch "$projName".go
touch input.txt

cat > "./$projName.go" << EOF
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    lines, err := readLines("./input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines, scanner.Err()
}
EOF
