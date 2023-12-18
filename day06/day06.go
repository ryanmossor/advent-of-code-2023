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

func partOne(lines []string) (error) {
    times := parseInput(lines[0])
    distances := parseInput(lines[1])

    var totals []int

    for i, raceTime := range times {
        currentRecord := distances[i]
        possibleWins := 0

        for timeHeld := 0; timeHeld < raceTime; timeHeld++ {
            dist := timeHeld * (raceTime - timeHeld)
            if dist > currentRecord {
                possibleWins += 1
            }
        }
        totals = append(totals, possibleWins)
    }

    p := 1
    for _, total := range totals {
        p *= total
    }

    fmt.Println("Part 1:", p)

    return nil
}

func partTwo(lines []string) (error) {
    time := parseInputPartTwo(lines[0])
    currentRecord := parseInputPartTwo(lines[1])

    possibleWins := 0

    for timeHeld := int64(0); timeHeld < time; timeHeld++ {
        dist := timeHeld * (time - timeHeld)
        if dist > currentRecord {
            possibleWins += 1
        }
    }

    fmt.Println("Part 2:", possibleWins)

    return nil
}

func parseInput(line string) ([]int) {
    var arr []int
    colonIndex := strings.Index(line, ":")
    splitStr := strings.Split(line[colonIndex+1:], " ")

    for _, x := range splitStr {
        if strings.Trim(x, " ") == "" {
            continue
        }

        converted, _ := strconv.Atoi(strings.Trim(x, " "))
        arr = append(arr, converted)
    }

    return arr
}

func parseInputPartTwo(line string) (int64) {
    colonIndex := strings.Index(line, ":")
    splitStr := strings.Split(line[colonIndex+1:], " ")

    numAsStr := ""

    for _, x := range splitStr {
        if strings.TrimSpace(x) == "" {
            continue
        }

        numAsStr += strings.TrimSpace(x)
    }

    converted, _ := strconv.Atoi(numAsStr)
    return int64(converted)
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
