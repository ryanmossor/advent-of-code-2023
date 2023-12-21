package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
    L string
    R string
}

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
    var regex = regexp.MustCompile(`(?P<node>[A-Z]{3}) = \((?P<left>[A-Z]{3}), (?P<right>[A-Z]{3})\)`)
    instructions := lines[0]
    nodeMap := make(map[string]Node)

    for _, line := range lines {
        match := regex.FindStringSubmatch(line)

        if len(match) == 0 {
            continue
        }

        node, left, right := match[1], match[2], match[3]
        n := Node { left, right }
        nodeMap[node] = n
    }

    currentNode := "AAA"
    steps := 0
    for !strings.EqualFold(currentNode, "ZZZ") {
        for _, ch := range instructions {
            if string(ch) == "L" {
                currentNode = nodeMap[currentNode].L
            } else {
                currentNode = nodeMap[currentNode].R
            }

            steps++

            if strings.EqualFold(currentNode, "ZZZ") {
                break
            }
        }
    }

    fmt.Println("Part 1:", steps)
    return nil
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
