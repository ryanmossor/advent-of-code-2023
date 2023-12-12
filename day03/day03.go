package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type part struct {
    numStr string
    partNumber int
    index int
    lineNumber int
}

type symbol struct {
    sym string
    index int
    lineNumber int
}

func main() {
    lines, err := readLines("./input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

    partOneResult, err := partOne(lines)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    fmt.Fprintf(os.Stdout, "Part 1: %d\n", partOneResult)

    partTwoResult, err := partTwo(lines)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    fmt.Fprintf(os.Stdout, "Part 2: %d\n", partTwoResult)

}

func partOne(lines []string) (int, error) {
    symbolsRegex := regexp.MustCompile(`[^0-9a-zA-Z.]`)
    var symbols []symbol
    var parts []part

    for lineNum, line := range lines {
        if matchedSymbols := symbolsRegex.FindAllStringIndex(line, -1); matchedSymbols != nil {
            for _, idx := range matchedSymbols {
                s := symbol { 
                    sym: line[idx[0]:idx[1]],
                    index: idx[0],
                    lineNumber: lineNum,
                }
                symbols = append(symbols, s)
            }
        }

        parts = append(parts, parseNumbers(line, lineNum)...)
    }

    total := 0
    for _, part := range parts {
        symLowerBound := part.index - 1
        symUpperBound := len(part.numStr) + part.index

        for _, symbol := range symbols {
            if symbol.lineNumber >= part.lineNumber - 1 && symbol.lineNumber <= part.lineNumber + 1 {
                if symbol.index >= symLowerBound && symbol.index <= symUpperBound {
                    total += part.partNumber
                }
            }
        }
    }

    return total, nil
}

func partTwo(lines []string) (int, error) {
    var symbols []symbol
    var parts []part

    for lineNum, line := range lines {
        for idx, char := range line {
            if char == '*' {
                s := symbol{
                	sym: string(char),
                	index: idx,
                	lineNumber: lineNum,
                }
                symbols = append(symbols, s)
            }
        }

        parts = append(parts, parseNumbers(line, lineNum)...)
    }

    total := 0
    for _, asterisk := range symbols {
        var adjacentPartNumbers []int

        for _, part := range parts {
            if part.lineNumber >= asterisk.lineNumber - 1 && part.lineNumber <= asterisk.lineNumber + 1 {
                lowerBound := part.index - 1
                upperBound := part.index + len(part.numStr)

                if asterisk.index >= lowerBound && asterisk.index <= upperBound {
                    adjacentPartNumbers = append(adjacentPartNumbers, part.partNumber)
                }
            }
        }

        if len(adjacentPartNumbers) == 2 {
            total += adjacentPartNumbers[0] * adjacentPartNumbers[1]
        }
    }

    return total, nil
}

func parseNumbers(line string, lineNum int) ([]part) {
    numbersRegex := regexp.MustCompile(`[0-9]*`)
    var parts []part

    if matchedNumbers := numbersRegex.FindAllStringIndex(line, -1); matchedNumbers != nil {
        for _, idx := range matchedNumbers {
            numAsStr := line[idx[0]:idx[1]]
            if numAsStr == "" {
                continue
            }

            num, err := strconv.Atoi(numAsStr) 
            if err != nil {
                panic(err)
            }

            p := part { 
                numStr: numAsStr,
                partNumber: num,
                index: idx[0],
                lineNumber: lineNum,
            }

            parts = append(parts, p)
        }
    }

    return parts
}
