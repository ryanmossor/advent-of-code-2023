package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type seedCategories struct {
    seed int
    soil int
    fertilizer int
    water int
    light int
    temperature int
    humidity int
    location int
}

func main() {
    lines, err := readLines("./input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

    err = partOne(lines)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func partOne(lines []string) (error) {
    var seeds []int
    var seedInfoArr []seedCategories
    var currentMap string

    seedsLine := lines[0]

    for _, seed := range strings.Split(seedsLine[strings.Index(seedsLine, ":") + 1:], " ") {
        if seed == "" {
            continue
        }

        seedAsInt, err := strconv.Atoi(strings.Trim(seed, " "))
        if err != nil {
            return err
        }

        seeds = append(seeds, seedAsInt)
    }

    for _, seed := range seeds {
        var seedInfo seedCategories
        seedInfo.seed = seed
        seedInfo.soil = seed

        skipToNextMap := false

        for _, line := range lines {
            if line == "" || strings.HasPrefix(line, "seeds:") {
                continue
            }

            if skipToNextMap || strings.HasSuffix(line, "map:") {
                currentMap = strings.Split(line, " ")[0]
                skipToNextMap = false
                continue
            }

            switch currentMap {
                case "seed-to-soil":
                    if soilNum := checkMap(line, seedInfo.seed, &skipToNextMap); soilNum > 0 {
                        seedInfo.soil = soilNum
                        seedInfo.fertilizer = soilNum
                        seedInfo.water = soilNum
                        seedInfo.light = soilNum
                        seedInfo.temperature = soilNum
                        seedInfo.humidity = soilNum
                        seedInfo.location = soilNum
                    }
                case "soil-to-fertilizer":
                    if fertilizerNum := checkMap(line, seedInfo.soil, &skipToNextMap); fertilizerNum > 0 {
                        seedInfo.fertilizer = fertilizerNum 
                        seedInfo.water = fertilizerNum 
                        seedInfo.light = fertilizerNum 
                        seedInfo.temperature = fertilizerNum 
                        seedInfo.humidity = fertilizerNum 
                        seedInfo.location = fertilizerNum 
                    }
                case "fertilizer-to-water":
                    if waterNum := checkMap(line, seedInfo.fertilizer, &skipToNextMap); waterNum > 0 {
                        seedInfo.water = waterNum 
                        seedInfo.light = waterNum 
                        seedInfo.temperature = waterNum 
                        seedInfo.humidity = waterNum 
                        seedInfo.location = waterNum 
                    }
                case "water-to-light":
                    if lightNum := checkMap(line, seedInfo.water, &skipToNextMap); lightNum > 0 {
                        seedInfo.light = lightNum 
                        seedInfo.temperature = lightNum 
                        seedInfo.humidity = lightNum 
                        seedInfo.location = lightNum 
                    }
                case "light-to-temperature":
                    if temperatureNum := checkMap(line, seedInfo.light, &skipToNextMap); temperatureNum > 0 {
                        seedInfo.temperature = temperatureNum 
                        seedInfo.humidity = temperatureNum 
                        seedInfo.location = temperatureNum 
                    }
                case "temperature-to-humidity":
                    if humidityNum := checkMap(line, seedInfo.temperature, &skipToNextMap); humidityNum > 0 {
                        seedInfo.humidity = humidityNum 
                        seedInfo.location = humidityNum 
                    }
                case "humidity-to-location":
                    if locationNum := checkMap(line, seedInfo.humidity, &skipToNextMap); locationNum > 0 {
                        seedInfo.location = locationNum 
                    }
            }
        }

        seedInfoArr = append(seedInfoArr, seedInfo)
    }

    lowest := -1
    for _, s := range seedInfoArr {
        if lowest == -1 || s.location < lowest {
            lowest = s.location
        }
    }

    fmt.Println("Part 1:", lowest)

    return nil
}

func checkMap(line string, seed int, matchFound *bool) (int) {
    nums := strings.Split(line, " ")

    destRangeStart, err := strconv.Atoi(nums[0])
    if err != nil {
        panic(err)
    }

    sourceRangeStart, err := strconv.Atoi(nums[1])
    if err != nil {
        panic(err)
    }

    rangeLength, err := strconv.Atoi(nums[2])
    if err != nil {
        panic(err)
    }

    if seed >= sourceRangeStart && seed <= sourceRangeStart + rangeLength {
        *matchFound = true
        return seed + (destRangeStart - sourceRangeStart)
    }

    return -1
}
