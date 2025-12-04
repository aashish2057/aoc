use std::fs::read_to_string;

pub fn part1() -> i32 {
    let mut grid: Vec<Vec<char>> = Vec::new();
    for line in read_to_string("./inputs/2025day4.txt").unwrap().lines() {
        grid.push(line.chars().collect());
    }

    forklift(grid)
}

// grid traversal
// for every position check all adjacent positions
// if adjacent position in grid and @ add to count
// if > 4 move on if < 4 ++
fn forklift(grid: Vec<Vec<char>>) -> i32 {
    let mut accessible_rolls = 0;

    // println!("{:?}", grid);

    for (row_num, row) in grid.iter().enumerate() {
        for (col_num, col) in row.iter().enumerate() {
            let mut adjacent_rolls = 0;

            if col == &'.' {
                continue;
            }
            // (-1,1)
            if row_num > 0 && col_num + 1 < row.len() {
                if grid[row_num - 1][col_num + 1] == '@' {
                    adjacent_rolls += 1;
                }
            }
            // (0,1)
            if col_num + 1 < row.len() {
                if grid[row_num][col_num + 1] == '@' {
                    adjacent_rolls += 1;
                }
            }
            // (1,1)
            if row_num + 1 < grid.len() && col_num + 1 < row.len() {
                if grid[row_num + 1][col_num + 1] == '@' {
                    adjacent_rolls += 1;
                }
            }
            // (-1,0)
            if row_num > 0 {
                if grid[row_num - 1][col_num] == '@' {
                    adjacent_rolls += 1;
                }
            }
            // (0,0)
            // if col == &'@' {
            //     adjacent_rolls += 1;
            // }
            // (1,0)
            if row_num + 1 < grid.len() {
                if grid[row_num + 1][col_num] == '@' {
                    adjacent_rolls += 1;
                }
            }
            // (-1,-1)
            if row_num > 0 && col_num > 0 {
                if grid[row_num - 1][col_num - 1] == '@' {
                    adjacent_rolls += 1;
                }
            }
            // (0,-1)
            if col_num > 0 {
                if grid[row_num][col_num - 1] == '@' {
                    adjacent_rolls += 1;
                }
            }
            // (1,-1)
            if row_num + 1 < grid.len() && col_num > 0 {
                if grid[row_num + 1][col_num - 1] == '@' {
                    adjacent_rolls += 1;
                }
            }

            if adjacent_rolls < 4 {
                accessible_rolls += 1;
            }

            // println!(
            //     "acessible: {}, adjacnet: {}, ({},{})",
            //     accessible_rolls, adjacent_rolls, row_num, col_num
            // );
        }
    }

    accessible_rolls
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_forklift() {
        let grid: Vec<Vec<char>> = vec![
            vec!['.', '.', '@', '@', '.', '@', '@', '@', '@', '.'],
            vec!['@', '@', '@', '.', '@', '.', '@', '.', '@', '@'],
            vec!['@', '@', '@', '@', '@', '.', '@', '.', '@', '@'],
            vec!['@', '.', '@', '@', '@', '@', '.', '.', '@', '.'],
            vec!['@', '@', '.', '@', '@', '@', '@', '.', '@', '@'],
            vec!['.', '@', '@', '@', '@', '@', '@', '@', '.', '@'],
            vec!['.', '@', '.', '@', '.', '@', '.', '@', '@', '@'],
            vec!['@', '.', '@', '@', '@', '.', '@', '@', '@', '@'],
            vec!['.', '@', '@', '@', '@', '@', '@', '@', '@', '.'],
            vec!['@', '.', '@', '.', '@', '@', '@', '.', '@', '.'],
        ];

        assert_eq!(forklift(grid), 13);
    }
}
