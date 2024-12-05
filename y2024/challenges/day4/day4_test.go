package day4

import (
    "reflect"
    "testing"
)

func TestProcessFile(t *testing.T) {
    filePath := "./test.txt"

    expectedMatrix := []string{
        "MMMSXXMASM",
        "MSAMXMSMSA",
        "AMXSXMAAMM",
        "MSAMASMSMX",
        "XMASAMXAMM",
        "XXAMMXXAMA",
        "SMSMSASXSS",
        "SAXAMASAAA",
        "MAMMMXMMMM",
        "MXMXAXMASX",
    }

    actualMatrix := processFile(filePath)

    if !reflect.DeepEqual(actualMatrix, expectedMatrix) {
        t.Errorf("Actual: %v, Expected: %v, when calling processFile with: %s", actualMatrix, expectedMatrix, filePath)
    }
}

func TestWordSearch(t *testing.T) {
    grid := []string{
        "MMMSXXMASM",
        "MSAMXMSMSA",
        "AMXSXMAAMM",
        "MSAMASMSMX",
        "XMASAMXAMM",
        "XXAMMXXAMA",
        "SMSMSASXSS",
        "SAXAMASAAA",
        "MAMMMXMMMM",
        "MXMXAXMASX",
    }
    expectedCount := 18

    actualCount := wordSearch(grid, 'X')

    if actualCount != expectedCount {
        t.Errorf("Actual: %d, Expected: %d, when calling wordSearch with: %v", actualCount, expectedCount, grid)
    }

    actualCount2 := wordSearch(grid, 'A')
    expectedCount2 := 9

    if actualCount2 != expectedCount2 {
        t.Errorf("Actual: %d, Expected: %d, when calling wordSearch with: %v", actualCount2, expectedCount2, grid)
    }
}

func TestSearchForXmas(t *testing.T) {
    grid := []string{
        "MMMSXXMASM",
        "MSAMXMSMSA",
        "AMXSXMAAMM",
        "MSAMASMSMX",
        "XMASAMXAMM",
        "XXAMMXXAMA",
        "SMSMSASXSS",
        "SAXAMASAAA",
        "MAMMMXMMMM",
        "MXMXAXMASX",
    }

    tests := []struct {
        row_pos int
        col_pos int
        expectedXmasCount int
    }{
        {0, 5, 1},
        {0, 4, 1},
        {1, 4, 1},
        {3, 9, 2},
        {4, 0, 1},
        {4, 6, 2},
        {5, 0, 1},
        {5, 6, 1},
        {9, 1, 1},
        {9, 3, 2},
        {9, 5, 3},
        {9, 9, 2},
    }


    for _, test := range tests {
        actualXmasCount := searchForXmas(test.row_pos, test.col_pos, grid)

        if test.expectedXmasCount != actualXmasCount {
            t.Errorf("Actual: %d, Expected: %d, when calling searchForXmas with: %d, %d, %v", actualXmasCount, test.expectedXmasCount,
                test.row_pos, test.col_pos, grid)
        }
    }
}

func TestSearchForXmas2(t *testing.T) {
    grid := []string{
        "MMMSXXMASM",
        "MSAMXMSMSA",
        "AMXSXMAAMM",
        "MSAMASMSMX",
        "XMASAMXAMM",
        "XXAMMXXAMA",
        "SMSMSASXSS",
        "SAXAMASAAA",
        "MAMMMXMMMM",
        "MXMXAXMASX",
    }

    tests := []struct {
        row_pos int
        col_pos int
        expectedXmasCount int
    }{
        {1, 2, 1},
        {2, 6, 1},
        {2, 7, 1},
        {3, 2, 1},
        {3, 4, 1},
        {7, 1, 1},
        {7, 3, 1},
        {7, 5, 1},
        {7, 7, 1},
    }

    for _, test := range tests {
        actualXmasCount := searchForXmas2(test.row_pos, test.col_pos, grid)

        if test.expectedXmasCount != actualXmasCount {
            t.Errorf("Actual: %d, Expected: %d, when calling searchForXmas2 with: %d, %d, %v", actualXmasCount, test.expectedXmasCount,
                test.row_pos, test.col_pos, grid)
        }
    }
}
