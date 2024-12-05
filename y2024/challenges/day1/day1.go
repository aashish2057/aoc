package day1

import (
    "bufio"
    "os"
    "sort"
    "strconv"
    "strings"

    "github.com/aashish2057/aoc/y2024/util"
)

func Part1() int {
    path := "./inputs/day1.txt"

    locations1, locations2 := processFile(path)

    return calculateTotalDistance(locations1, locations2)
}

func Part2() int {
    filePath := "./inputs/day1.txt"

    locations1, locations2 := processFile(filePath)

    mapLocations2 := getMapOccuranceOfLocationIds(locations2)

    return calculateSimilarityScore(locations1, mapLocations2)
}

func processFile(path string) ([]int, []int) {
    var locations1 []int
    var locations2 []int

    file, err := os.Open(path)
    defer file.Close()

    util.CheckError(err)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        locationIds := strings.Split(scanner.Text(), "   ")
        num1, err := strconv.Atoi(locationIds[0])
        util.CheckError(err)

        num2, err := strconv.Atoi(locationIds[1])
        util.CheckError(err)

        locations1 = append(locations1, num1)
        locations2 = append(locations2, num2)
    }

    sort.Ints(locations1)
    sort.Ints(locations2)

    return locations1, locations2
}

func calculateTotalDistance(locations1 []int, locations2 []int) int {
    totalDistance := 0

    for index := 0; index < len(locations1); index++ {
        totalDistance += util.Abs(locations1[index] - locations2[index])
    }

    return totalDistance
}

func getMapOccuranceOfLocationIds(locations []int) map[int]int {
    locationIdOccurance := make(map[int]int)

    for _, locationId := range locations {
        if _, exists := locationIdOccurance[locationId]; !exists {
            locationIdOccurance[locationId] = 0
        }
        locationIdOccurance[locationId] += 1
    }
    return locationIdOccurance
}

func calculateSimilarityScore(locations1 []int, mapLocations2 map[int]int) int {
    similarityScore := 0

    for _, num := range locations1 {
        similarityScore += num * mapLocations2[num]
    }

    return similarityScore
}
