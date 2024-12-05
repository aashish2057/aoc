package day2

import (
    "testing"
    "reflect"
)

func TestProcessFile(t *testing.T) {
    filePath := "./test.txt"

    expectedReports := [][]int{
        {7, 6, 4, 2, 1},
        {1, 2, 7, 8, 9},
        {9, 7, 6, 2, 1},
        {1, 3, 2, 4, 5},
        {8, 6, 4, 4, 1},
        {1, 3, 6, 7, 9},
    }

    actualReports := processFile(filePath)

    if !reflect.DeepEqual(actualReports, expectedReports) {
        t.Errorf("Expected: %v, Actual: %v, when calling processFile with: %s", expectedReports, actualReports, filePath)
    }

}

func TestIsReportSafe(t *testing.T) {
    tests := []struct {
        report []int
        expectedSafe bool
    }{
        {[]int{7, 6, 4, 2, 1}, true},
        {[]int{1, 2, 7, 8, 9}, false},
        {[]int{9, 7, 6, 2, 1}, false},
        {[]int{9, 7, 6, 2, 10}, false},
        {[]int{1, 3, 2, 4, 5}, false},
        {[]int{8, 6, 4, 4, 1}, false},
        {[]int{1, 3, 6, 7, 9}, true},
        {[]int{1, 1, 6, 7, 9}, false},
    }

    for _, test := range tests {
        actualSafe := isReportSafe(test.report)

        if actualSafe != test.expectedSafe {
            t.Errorf("Expected: %v, Actual: %v, when calling isReportSafe with: %v", test.expectedSafe, actualSafe, test.report)
        }
    }
}


func TestCalculateTotalSafeReports(t *testing.T) {
    reports := [][]int {
        {7, 6, 4, 2, 1},
        {1, 2, 7, 8, 9},
        {9, 7, 6, 2, 1},
        {9, 7, 6, 2, 10},
        {1, 3, 2, 4, 5},
        {8, 6, 4, 4, 1},
        {1, 3, 6, 7, 9},
        {1, 1, 6, 7, 9},
        {1, 2, 3, 2},
    }
    expectedSafeCount := 5

    actualSafeCount := calculateTotalSafeReportsWithTolerating(reports)

    if actualSafeCount != expectedSafeCount {
        t.Errorf("Expected: %d, Actual: %d, when calling calculateTotalSafeReports with: %v", expectedSafeCount,
            actualSafeCount, reports)
    }
}
