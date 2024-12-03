package day2


// What is the problem?
//  Determine if an report is safe
// A report consists of a list of numbers called levels
// Each line is a report contains the levels
// For a level to be safe it must meet the follow criteria
// All numbers must be increasing, or decrease and adjacent levels must differ by atleats one and at most 3

// Sub problems
// Read the file in line by line
// For each line determine if the levels are "safe"
//      They are always increase or decrease and adjacent levels differ by atleast one or at most 3
// Count total number of reports that are "safe"

import (
    "os"
    "bufio"
    "strings"
    "strconv"
    "fmt"

    "github.com/aashish2057/aoc/y2024/util"
)

func Part1() int {
    file, err := os.Open("./inputs/day2.txt")
    util.CheckError(err)
    defer file.Close()

    totalSafeReports := 0

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        report := strings.Fields(scanner.Text())
        newReport := convertReportToInt(report)

        if isReportSafe(newReport) {
            totalSafeReports += 1
        }
    }

    return totalSafeReports
}

func Part2() int {
    file, err := os.Open("./inputs/day2.txt")
    util.CheckError(err)
    defer file.Close()

    totalSafeReports := 0

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        report := strings.Fields(scanner.Text())
        newReport := convertReportToInt(report)

        if !isReportSafe(newReport) {

            for i := 0; i < len(report); i++ {
                dampenedReport := append(newReport[:i], newReport[i+1:]...)

                if isReportSafe(dampenedReport) {
                    fmt.Println("every here?")
                    totalSafeReports += 1
                }
            }
        } else {
            totalSafeReports += 1
        }
    }

    return totalSafeReports
}

func isReportSafe(report []int) bool {
    isSafe := true
    increasing := report[1] >= report[0]

    for i := 1; i < len(report); i++ {
        diff := report[i] - report[i-1]

        if (increasing && diff < 0) || (!increasing && diff > 0) || util.Abs(diff) > 3 || diff == 0 {
            isSafe = false
            break
        }
    }
    return isSafe
}

func convertReportToInt(report []string) []int {

    var intReport []int

    for _, level := range report {
        num, err := strconv.Atoi(level)
        util.CheckError(err)

        intReport = append(intReport, num)
    }

    return intReport
}

