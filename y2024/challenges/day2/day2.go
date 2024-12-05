package day2

import (
    "bufio"
    "strconv"
    "os"
    "strings"

    "github.com/aashish2057/aoc/y2024/util"
)

func Part1() int {
    filePath := "./inputs/day2.txt"

    reports := processFile(filePath)

    return calculateTotalSafeReports(reports)
}

func Part2() int {
    filePath := "./inputs/day2.txt"

    reports := processFile(filePath)

    return calculateTotalSafeReportsWithTolerating(reports)
}

func processFile(path string) [][]int {
    var reports [][]int

    file, err := os.Open(path)
    defer file.Close()
    util.CheckError(err)

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        stringReport := strings.Split(scanner.Text(), " ")
        var report []int

        for _, char := range stringReport {
            num, err := strconv.Atoi(char)
            util.CheckError(err)

            report = append(report, num)
        }

        reports = append(reports, report)
    }

    return reports
}

func calculateTotalSafeReports(reports [][]int) int {
    totalSafeReports := 0

    for _, report := range reports {
        if isReportSafe(report) {
            totalSafeReports += 1
        }
    }

    return totalSafeReports
}

func calculateTotalSafeReportsWithTolerating(reports [][]int) int {
    totalSafeReports := 0

    for _, report := range reports {
        if isReportSafe(report) {
            totalSafeReports += 1
        } else {
            for index, _ := range report {
                cleanedReport := append([]int{}, report[:index]...)
		        cleanedReport = append(cleanedReport, report[index+1:]...)

                if isReportSafe(cleanedReport) {
                    totalSafeReports += 1
                    break
                }
	        }
        }
    }

    return totalSafeReports
}

func isReportSafe(report []int) bool {
    safe := true

    if !isIncreasing(report) && !isDecreasing(report) {
        safe = false
    }

    for index := 1; index < len(report); index ++ {
        diff := util.Abs(report[index] - report[index - 1])

        if diff < 1 || diff > 3 {
            safe = false
            break
        }
    }

    return safe
}

func isIncreasing(report []int) bool {
    increasing := true

    for index := 1; index < len(report); index ++ {
        if report[index] < report[index-1] {
            increasing = false
        }
    }

    return increasing

}

func isDecreasing(report []int) bool {
    decreasing := true

    for index := 1; index < len(report); index ++ {
        if report[index] > report[index-1] {
            decreasing = false
        }
    }

    return decreasing
}
