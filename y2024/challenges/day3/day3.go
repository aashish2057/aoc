package day3

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "regexp"

    "fmt"

    "github.com/aashish2057/aoc/y2024/util"
)

func Part1() int {
    filePath := "./inputs/day3.txt"

    memories := processFile(filePath)
    totalProduct := 0

    for _, memory := range memories {

        validSubStrings := findAllValidSubStrings(memory, `mul\(\d{1,3},\d{1,3}\)`)

        for _, subString := range validSubStrings {
            totalProduct += calculateProduct(subString)
        }

    }

    return totalProduct
}

func Part2() int {
    filePath := "./inputs/day3.txt"

    memories := processFile(filePath)
    totalProduct := 0

    for _, memory := range memories {

        validSubStrings := findAllValidSubStrings(memory, `(mul\(\d{1,3},\d{1,3}\))|(don't\(\))|(do\(\))`)

        cleanedSubStrings := removeDisabledInstructions(validSubStrings)

        for _, subString := range cleanedSubStrings {
            totalProduct += calculateProduct(subString)
        }
    }

    return totalProduct
}

func processFile(path string) []string {
    var memories []string

    file, err := os.Open(path)
    util.CheckError(err)

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        memories = append(memories, scanner.Text())
    }

    return memories
}

func findAllValidSubStrings(memory string, expression string) []string {
    regex := regexp.MustCompile(expression)

    validSubStrings := regex.FindAllString(memory, -1)

    return validSubStrings
}

func removeDisabledInstructions(subStrings []string) []string {
    disabled := false

    var cleanedSubStrings []string
    var debug []string

    for _, subString := range subStrings {
        if subString == "don't()" {
            // fmt.Println(index, subString)
            disabled = true
        } else if subString == "do()" {
            // fmt.Println(index, subString)
            disabled = false
        } else {
            if !disabled {
                // fmt.Println(index, "included", subString)
                cleanedSubStrings = append(cleanedSubStrings, subString)
            } else {
                // fmt.Println(index, "disabled", subString)
                debug = append(debug, subString)
            }
        }
    }
    return cleanedSubStrings
}

func calculateProduct(substring string) int {
    cleanedSubstring := substring[4: len(substring) - 1]

    numbers := strings.Split(cleanedSubstring, ",")

    num1, err := strconv.Atoi(numbers[0])
    util.CheckError(err)

    num2, err := strconv.Atoi(numbers[1])
    util.CheckError(err)

    fmt.Println(substring, num1 * num2)
    return num1 * num2
}
