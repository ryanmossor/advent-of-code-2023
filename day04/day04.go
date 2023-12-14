package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
    lines, err := readLines("./input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

    partOneResult, err := partOne(lines)
    if err != nil {
        panic(err)
    }
    fmt.Println("Part 1:", partOneResult)

    partTwoResult, err := partTwo(lines)
    if err != nil {
        panic(err)
    }
    fmt.Println("Part 2:", partTwoResult)
}

func partOne(lines []string) (float64, error) {
    var total float64 = 0

    for _, line := range lines {
        colonIndex := strings.Index(line, ":")
        winningNums := strings.Split(strings.Split(strings.Trim(line[colonIndex + 1:], " "), "|")[0], " ")
        ourNums := strings.Trim(strings.Split(line, "|")[1], " ")

        var matches float64 = 0
        for _, num := range strings.Split(ourNums, " ") {
            if num == "" {
                continue
            }

            if slices.Contains(winningNums, num) {
                matches++
            }
        }
        
        if matches > 0 {
            total += math.Pow(2, matches - 1)
        }
    }

    return total, nil
}

func partTwo(lines []string) (int, error) {
    cardCopies := make(map[int]int)
    for i := range lines {
        cardCopies[i] += 1
    }

    for i, line := range lines {
        colonIndex := strings.Index(line, ":")
        winningNums := strings.Split(strings.Split(strings.Trim(line[colonIndex + 1:], " "), "|")[0], " ")
        ourNums := strings.Trim(strings.Split(line, "|")[1], " ")

        matches := 0
        for _, num := range strings.Split(ourNums, " ") {
            if num == "" {
                continue
            }

            if slices.Contains(winningNums, num) {
                matches++
            }
        }
        
        copiesOfCurrentCard := cardCopies[i]

        if matches > 0 {
            for idx := i + 1; idx <= i + matches; idx++ {
                cardCopies[idx] += 1 * copiesOfCurrentCard
            }
        }
    }

    total := 0
    for _, val := range cardCopies {
        total += val
    }

    return total, nil
}
