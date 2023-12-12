#! /usr/bin/env bash

if [[ -z "$1" ]]; then
    echo "No project name provided"
    exit 1
fi

projName="$1"

mkdir -p "$projName"
cp ./utils.go "$projName"

cd "$projName"

touch "$projName".go
touch input.txt

cat > "./$projName.go" << EOF
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    lines, err := readLines("./input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

EOF
