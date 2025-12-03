use std::fs::read_to_string;

pub fn part1() -> i64 {
    let mut sum = 0;

    for bank in read_to_string("./inputs/day3.txt").unwrap().lines() {
        sum += largest_joltage(&bank);
    }

    sum
}

fn largest_joltage(bank: &str) -> i64 {
    let mut res: [i64; 2] = [0; 2];

    let mut ptr2 = 1;

    res[0] = bank[0..ptr2].parse().unwrap();

    while ptr2 < bank.len() {
        let num: i64 = bank[ptr2..ptr2 + 1].parse().unwrap();

        // Larger than first digit and not last joltage in bank
        if num > res[0] && ptr2 != bank.len() - 1 {
            res[0] = num;
            res[1] = 0;
        }
        // Larger than first digit and last joltage in bank
        else if num > res[0] && ptr2 == bank.len() - 1 {
            res[1] = num;
        }
        // Smaller than equal first digit and greater than 2nd digit
        else if num <= res[0] && num >= res[1] {
            res[1] = num;
        }

        ptr2 += 1
    }

    (res[0] * 10) + res[1]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_largest_joltage() {
        assert_eq!(largest_joltage("987654321111111"), 98);
        assert_eq!(largest_joltage("811111111111119"), 89);
        assert_eq!(largest_joltage("234234234234278"), 78);
        assert_eq!(largest_joltage("818181911112111"), 92);
        assert_eq!(
            largest_joltage(
                "6483266694748235893324353634344523834567333718239477213324541343624714732212276727733744455653544463"
            ),
            99
        );
    }

    #[test]
    fn test_total_joltage() {
        let mut sum = 0;

        sum += largest_joltage("987654321111111");
        sum += largest_joltage("811111111111119");
        sum += largest_joltage("234234234234278");
        sum += largest_joltage("818181911112111");

        assert_eq!(sum, 357);
    }
}
