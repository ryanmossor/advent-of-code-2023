package main

import (
	"bufio"
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

    partOne(lines)
    partTwo(lines)
}

func partOne(lines []string) error {
    total := 0

    for _, line := range lines {
        if line == "" {
            break
        }

        sequences := parseSequences(line)
        for i := len(sequences) - 1; i > 0; i-- {
            lastVal1 := sequences[i][len(sequences[i]) - 1]
            lastVal2 := sequences[i - 1][len(sequences[i - 1]) - 1]

            sequences[i - 1] = append(sequences[i - 1], lastVal1 + lastVal2)
        }

        total += sequences[0][len(sequences[0]) - 1]
    }

    fmt.Println("Part 1:", total)
    return nil
}

func partTwo(lines []string) error {
    total := 0

    for _, line := range lines {
        if line == "" {
            break
        }

        sequences := parseSequences(line)
        for i := len(sequences) - 1; i > 0; i-- {
            lastVal1 := sequences[i][0]
            lastVal2 := sequences[i - 1][0]

            sequences[i - 1] = append([]int { lastVal2 - lastVal1 }, sequences[i - 1]...)
        }

        total += sequences[0][0]
    }

    fmt.Println("Part 2:", total)
    return nil
}

func parseSequences(line string) [][]int {
    split := strings.Split(line, " ")
    startingHistory := []int {}

    for _, item := range split {
        num, _ := strconv.Atoi(item)
        startingHistory = append(startingHistory, num)
    }

    sequences := [][]int {}
    sequences = append(sequences, startingHistory)

    for !allValuesZero(sequences[len(sequences) - 1]) {
        currentArr := sequences[len(sequences) - 1]
        lineResults := []int {}

        for i := 0; i < len(currentArr) - 1; i++ {
            a, b := currentArr[i], currentArr[i + 1]
            lineResults = append(lineResults, b - a)
        }

        sequences = append(sequences, lineResults)
    }

    return sequences
}

func allValuesZero(ints []int) bool {
    if len(ints) == 0 {
        return false
    }

    for _, n := range ints {
        if n != 0 {
            return false
        }
    }

    return true
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
