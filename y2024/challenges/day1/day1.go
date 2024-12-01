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
    file, error := os.Open("./inputs/day1.txt")
    util.CheckError(error)
    historian1, historian2 := getHistorianLists(file)
    defer file.Close()

    return getTotalDistance(historian1, historian2)
}

func Part2() int {
    file, error := os.Open("./inputs/day1.txt")
    util.CheckError(error)
    historian1, historian2 := getHistorianLists(file)

    numberOfOccurances := getMapOfNumberOccurances(historian2)

    return getSimilarityScore(historian1, numberOfOccurances)
}

func getHistorianLists(file *os.File) ([]int, []int) {
    scanner := bufio.NewScanner(file)

    var historian1 []int
    var historian2 []int

    for scanner.Scan() {
        numbers := strings.Fields(scanner.Text())

        num1, err := strconv.Atoi(numbers[0])
        util.CheckError(err)

        num2, err := strconv.Atoi(numbers[1])
        util.CheckError(err)

        historian1 = append(historian1, num1)
        historian2 = append(historian2, num2)
    }

    sort.Sort(sort.IntSlice(historian1))
    sort.Sort(sort.IntSlice(historian2))

    return historian1, historian2
}

func getTotalDistance(historian1 []int, historian2 []int) int {
    var totalDistance int

    for i, _ := range historian1 {
        distance := historian1[i] - historian2[i]

        if distance < 0 {
            distance = -distance
        }

        totalDistance += distance
    }

    return totalDistance
}

func getMapOfNumberOccurances(nums []int) map[int]int {
    numberOfOccurances := make(map[int]int)

    for _, number := range nums {

        if _, exists := numberOfOccurances[number]; !exists {
            numberOfOccurances[number] = 0
        }
        numberOfOccurances[number] += 1
    }

    return numberOfOccurances
}

func getSimilarityScore(historian1 []int, numberOfOccurances map[int]int) int {
    var totalSimilarityScore = 0

    for _, locationId := range historian1 {
       totalSimilarityScore += locationId * numberOfOccurances[locationId]
    }

    return totalSimilarityScore
}

