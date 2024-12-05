package day4

import (
    "os"
    "bufio"
    // "fmt"

    "github.com/aashish2057/aoc/y2024/util"
)

func Part1() int {
    path := "./inputs/day4.txt"

    grid := processFile(path)

    return wordSearch(grid, 'X')
}

func Part2() int {
    path := "./inputs/day4.txt"

    grid := processFile(path)

    return wordSearch(grid, 'A')
}

func processFile(path string) []string {
    file, err := os.Open(path)
    util.CheckError(err)

    var grid []string

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        grid = append(grid, scanner.Text())
    }

    return grid
}

func wordSearch(grid []string, initalCharacter byte) int {
    foundWords := 0

    for row_pos, row := range grid {
        for col_pos, _ := range row {
            if grid[row_pos][col_pos] == initalCharacter {
                count := 0
                if initalCharacter == 'X' {
                    count = searchForXmas(row_pos, col_pos, grid)
                } else {
                    count = searchForXmas2(row_pos, col_pos, grid)
                }
                foundWords += count
            }
        }
    }

    return foundWords
}

func searchForXmas(row_pos int, col_pos int, grid []string) int {
    xmasCount := 0

    moves := []struct {
        x int
        y int
    }{
        {0, 1},
        {1, 0},
        {0, -1},
        {-1, 0},
        {1, -1},
        {-1, 1},
        {1, 1},
        {-1, -1},
    }

    searchWord := "MAS"
    word_ptr := 0

    x_length := len(grid)
    y_length := len(grid[0])

    curr_row := row_pos
    curr_col := col_pos

    for _, pos := range moves {
        curr_row += pos.x
        curr_col += pos.y

        inBounds := (curr_row >= 0 && curr_row < x_length) && (curr_col >= 0 && curr_col < y_length)

        for inBounds && word_ptr < len(searchWord) && grid[curr_row][curr_col] == searchWord[word_ptr] {
            // fmt.Println(pos.x, pos.y, curr_row, curr_col, string(grid[curr_row][curr_col]), string(searchWord[word_ptr]))

            curr_row += pos.x
            curr_col += pos.y
            word_ptr += 1

            inBounds = (curr_row >= 0 && curr_row < x_length) && (curr_col >= 0 && curr_col < y_length)
        }

        if word_ptr == len(searchWord) {
            xmasCount += 1
        }

        curr_row = row_pos
        curr_col = col_pos
        word_ptr = 0
    }

    return xmasCount
}

func searchForXmas2(row_pos int, col_pos int, grid []string) int {
    x_length := len(grid)
    y_length := len(grid[0])

    // fmt.Println("DEBUGGER", row_pos, col_pos)

    if (row_pos - 1 >= 0 && row_pos + 1 < x_length) && (col_pos - 1 >= 0 && col_pos + 1 < y_length) {
        if (grid[row_pos + 1][col_pos + 1] == 'S') && (grid[row_pos + 1][col_pos - 1] == 'S') {
            if (grid[row_pos - 1][col_pos + 1] == 'M') && (grid[row_pos - 1][col_pos - 1] == 'M') {
                return 1
            }
        }

        if (grid[row_pos + 1][col_pos + 1] == 'M') && (grid[row_pos + 1][col_pos - 1] == 'M') {
            if (grid[row_pos - 1][col_pos + 1] == 'S') && (grid[row_pos - 1][col_pos - 1] == 'S') {
                return 1
            }
        }

        if (grid[row_pos + 1][col_pos + 1] == 'M') && (grid[row_pos - 1][col_pos + 1] == 'M') {
            if (grid[row_pos - 1][col_pos - 1] == 'S') && (grid[row_pos + 1][col_pos - 1] == 'S') {
                return 1
            }
        }

        if (grid[row_pos + 1][col_pos + 1] == 'S') && (grid[row_pos - 1][col_pos + 1] == 'S') {
            if (grid[row_pos - 1][col_pos - 1] == 'M') && (grid[row_pos + 1][col_pos - 1] == 'M') {
                return 1
            }
        }
    }

    return 0
}

