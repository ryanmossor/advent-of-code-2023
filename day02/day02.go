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

func partOne(input []string) (int, error) {
    totalRedCubes := 12
    totalGreenCubes := 13
    totalBlueCubes := 14

    var validGames []int

    for _, line := range input {
        if line == "" {
            break
        }

        gameNumIndex := strings.Index(line, " ") + 1
        colonIndex := strings.Index(line, ":")
        gameNum, err := strconv.Atoi(line[gameNumIndex:colonIndex])
        if err != nil {
            panic(err)
        }

        gameIsValid := true
        for _, round := range strings.Split(line[colonIndex+1:], ";") {
            redCubes := 0
            greenCubes := 0
            blueCubes := 0

            for _, cubes := range strings.Split(round, ",") {
                x := strings.Split(strings.Trim(cubes, " "), " ")

                count, err := strconv.Atoi(x[0])
                if err != nil {
                    panic(err)
                }

                switch color := x[1]; color {
                case "red":
                    redCubes = count
                case "green":
                    greenCubes = count
                case "blue":
                    blueCubes = count
                }

                if redCubes > totalRedCubes || greenCubes > totalGreenCubes || blueCubes > totalBlueCubes {
                    gameIsValid = false;
                    continue
                }
            }

        }

        if gameIsValid {
            validGames = append(validGames, gameNum)
        }
    }

    result, err := sumSlice(validGames)
    if err != nil {
        panic(err)
    }

    return result, nil
}

func partTwo(input []string) (int, error) {
    var powers []int

    for _, line := range input {
        if line == "" {
            break
        }

        colonIndex := strings.Index(line, ":")

        minRedCubes := 0
        minGreenCubes := 0
        minBlueCubes := 0

        for _, round := range strings.Split(line[colonIndex+1:], ";") {
            for _, cubes := range strings.Split(round, ",") {
                x := strings.Split(strings.Trim(cubes, " "), " ")

                count, err := strconv.Atoi(x[0])
                if err != nil {
                    panic(err)
                }

                switch color := x[1]; color {
                case "red":
                    if count > minRedCubes {
                        minRedCubes = count
                    }
                case "green":
                    if count > minGreenCubes {
                        minGreenCubes = count
                    }
                case "blue":
                    if count > minBlueCubes {
                        minBlueCubes = count
                    }
                }
            }
        }

        cubeSetPower := minRedCubes * minGreenCubes * minBlueCubes
        powers = append(powers, cubeSetPower)
    }

    result, err := sumSlice(powers)
    if err != nil {
        panic(err)
    }
    
    return result, nil
}
