package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
    cards string
    bid int
}

func main() {
    handValues := []string {
        "high-card",
        "one-pair",
        "two-pair",
        "three-of-a-kind",
        "full-house",
        "four-of-a-kind",
        "five-of-a-kind",
    }

    lines, err := readLines("./input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

    partOne(lines, handValues)
    partTwo(lines, handValues)
}

func partOne(lines []string, handValues []string) error {
    handMap := make(map[string][]Hand)

    for _, val := range handValues {
        handMap[val] = []Hand{}
    }

    for _, line := range lines {
        if line == "" {
            break
        }

        currentHandMap, hand := parseHand(line, false)

        handType := getHandType(currentHandMap)
        handMap[handType] = append(handMap[handType], hand)
    }

    cardValues := map[string]int {
        "2": 2,
        "3": 3,
        "4": 4,
        "5": 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9,
        "T": 10,
        "J": 11,
        "Q": 12,
        "K": 13,
        "A": 14,
    }

    for _, hand := range handMap {
        sort.Slice(hand, func(a, b int) bool {
            for i := 0; i < 5; i++ {
                cardA := string(hand[a].cards[i])
                cardB := string(hand[b].cards[i])
                if cardValues[cardA] == cardValues[cardB] {
                    continue 
                } else if cardValues[cardA] < cardValues[cardB] {
                    return true
                } else {
                    return false
                }
            }
            return false
        })
    }

    var results []int

    for _, val := range handValues {
        for _, hand := range handMap[val] {
            results = append(results, hand.bid)
        }
    }

    total := calcTotal(results)
    fmt.Println("Part 1:", total)

    return nil
} 

func partTwo(lines []string, handValues []string) error {
    handMap := make(map[string][]Hand)

    for _, val := range handValues {
        handMap[val] = []Hand{}
    }

    for _, line := range lines {
        if line == "" {
            break
        }

        currentHandMap, hand := parseHand(line, true)

        var highest string
        for k, v := range currentHandMap {
            if strings.EqualFold(highest, "") {
                highest = k
            } else if v > currentHandMap[highest] {
                highest = k
            }
        }
        
        jokerCount := strings.Count(hand.cards, "J")
        currentHandMap[highest] = currentHandMap[highest] + jokerCount

        handType := getHandType(currentHandMap)
        handMap[handType] = append(handMap[handType], hand)
    }

    cardValues := map[string]int {
        "2": 2,
        "3": 3,
        "4": 4,
        "5": 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9,
        "T": 10,
        "J": 1,
        "Q": 12,
        "K": 13,
        "A": 14,
    }

    for _, hand := range handMap {
        sort.Slice(hand, func(a, b int) bool {
            for i := 0; i < 5; i++ {
                cardA := string(hand[a].cards[i])
                cardB := string(hand[b].cards[i])
                if cardValues[cardA] == cardValues[cardB] {
                    continue 
                } else if cardValues[cardA] < cardValues[cardB] {
                    return true
                } else {
                    return false
                }
            }
            return false
        })
    }

    var results []int

    for _, val := range handValues {
        for _, hand := range handMap[val] {
            results = append(results, hand.bid)
        }
    }

    total := calcTotal(results)
    fmt.Println("Part 2:", total)

    return nil
}

func parseHand(line string, jokers bool) (handMap map[string]int, hand Hand) {
    cards := strings.Split(line, " ")[0]
    bid, _ := strconv.Atoi(strings.Split(line, " ")[1])

    currentHandMap := make(map[string]int)

    for _, ch := range cards {
        card := string(ch)
        if jokers && strings.EqualFold(card, "J") {
            continue
        }

        if _, ok := currentHandMap[card]; ok {
            currentHandMap[card]++
        } else {
            currentHandMap[card] = 1
        }
    }

    h := Hand { cards: cards, bid: bid }
    return currentHandMap, h
}

func getHandType(handMap map[string]int) string {
    var handType string
    switch len(handMap) {
        case 1:
            handType = "five-of-a-kind"
        case 2:
            if isOfAKind(4, handMap) {
                handType = "four-of-a-kind"
            } else {
                handType = "full-house"
            }
        case 3:
            if isOfAKind(3, handMap) {
                handType = "three-of-a-kind"
            } else {
                handType = "two-pair"
            }
        case 4:
            handType = "one-pair"
        case 5:
            handType = "high-card"
    }

    return handType
}

func isOfAKind(n int, currentHandMap map[string]int) (bool) {
    for _, val := range currentHandMap {
        if val == n {
            return true
        }
    }

    return false
}

func calcTotal(results []int) int {
    total := 0
    for i := 0; i < len(results); i++ {
        bid := results[i]
        rank := i + 1
        winnings := bid * rank
        total += winnings
    }

    return total
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
