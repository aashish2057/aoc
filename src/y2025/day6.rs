use std::fs::read_to_string;

pub fn part1() -> i64 {
    let mut homework: Vec<Vec<String>> =
        vec![Vec::new(); get_number_of_columns("./inputs/2025day6.txt")];

    parse("./inputs/2025day6.txt", &mut homework);

    finish_homework(homework)
}

fn get_number_of_columns(file_path: &str) -> usize {
    let input = read_to_string(file_path).expect("File not found");

    let mut lines = input.lines().peekable();

    let first_line = lines.peek().expect("Empty file");
    first_line.split_whitespace().count()
}

fn parse(file_path: &str, homework: &mut Vec<Vec<String>>) {
    let input = read_to_string(file_path).expect("File not found");

    for line in input.lines() {
        let numbers: Vec<String> = line.split_whitespace().map(|s| s.to_string()).collect();

        for (index, num) in numbers.iter().enumerate() {
            homework[index].push(num.to_string());
        }
    }
}

fn finish_homework(homework: Vec<Vec<String>>) -> i64 {
    let mut grand_total = 0;

    for problem in homework {
        let operator = problem.last().expect("Empty problem?");
        let total: i64;

        let numbers: Vec<i64> = problem
            .iter()
            .take(problem.len() - 1)
            .map(|num| num.parse().expect("Unable to parse"))
            .collect();

        match operator.as_str() {
            "+" => total = numbers.iter().sum(),
            "*" => total = numbers.iter().product(),
            _ => panic!("unexpected input"),
        }
        grand_total += total;
    }

    grand_total
}

fn finish_homework_2(homework: Vec<Vec<String>>) -> i64 {
    let mut grand_total = 0;

    for problem in homework {
        let operator = problem.last().expect("Empty problem?");

        let total: i64;

        let max_length = problem
            .iter()
            .take(problem.len() - 1)
            .map(|num| num.len())
            .max()
            .unwrap();

        let mut numbers: Vec<Vec<char>> = vec![Vec::new(); max_length];

        for i in 0..max_length {
            for j in 0..problem.len() - 1 {
                if i < problem[j].len() {
                    match problem[j].chars().nth(problem[j].len() - 1 - i) {
                        Some(c) => numbers[i].push(c),
                        _ => (),
                    }
                }
            }
        }

        println!("{:?}", numbers);

        let nums: Vec<i64> = numbers
            .iter()
            .map(|num| {
                let s: String = num.iter().collect();
                s.parse().unwrap()
            })
            .collect();

        match operator.as_str() {
            "+" => total = nums.iter().sum(),
            "*" => total = nums.iter().product(),
            _ => panic!("unexpected input"),
        }
        grand_total += total;
    }

    grand_total
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_get_number_of_columns() {
        assert_eq!(get_number_of_columns("./inputs/2025day6test.txt"), 4);
    }

    #[test]
    fn test_parse() {
        let mut homework: Vec<Vec<String>> =
            vec![Vec::new(); get_number_of_columns("./inputs/2025day6test.txt")];

        parse("./inputs/2025day6test.txt", &mut homework);

        assert_eq!(homework.len(), 4);
        assert_eq!(homework[0], vec!["123", "45", "6", "*"]);
        assert_eq!(homework[1], vec!["328", "64", "98", "+"]);
        assert_eq!(homework[2], vec!["51", "387", "215", "*"]);
        assert_eq!(homework[3], vec!["64", "23", "314", "+"]);
    }

    #[test]
    fn test_finish_homework() {
        let mut homework: Vec<Vec<String>> =
            vec![Vec::new(); get_number_of_columns("./inputs/2025day6test.txt")];

        parse("./inputs/2025day6test.txt", &mut homework);

        assert_eq!(finish_homework(homework), 4277556);
    }

    #[test]
    fn test_finish_homework_2() {
        let mut homework: Vec<Vec<String>> =
            vec![Vec::new(); get_number_of_columns("./inputs/2025day6test.txt")];

        parse("./inputs/2025day6test.txt", &mut homework);

        assert_eq!(finish_homework_2(homework), 3263827);
    }
}
