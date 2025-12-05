use std::fs::read_to_string;

pub fn part1() -> i64 {
    let input = read_to_string("./inputs/2025day5.txt").unwrap();

    let (new_id_ranges, fresh_ids) = input.split_once("\n\n").unwrap();

    let mut fresh_id_ranges: Vec<(i64, i64)> = Vec::new();

    for new_id_range in new_id_ranges.lines() {
        add_fresh_ids(new_id_range, &mut fresh_id_ranges);
    }

    let mut fresh_ingredients_count = 0;

    for fresh_id in fresh_ids.lines() {
        if is_fresh(fresh_id, &fresh_id_ranges) {
            fresh_ingredients_count += 1;
        }
    }

    fresh_ingredients_count
}

pub fn part2() -> i64 {
    let input = read_to_string("./inputs/2025day5.txt").unwrap();

    let (new_id_ranges, _) = input.split_once("\n\n").unwrap();

    let mut fresh_id_ranges: Vec<(i64, i64)> = Vec::new();

    for new_id_range in new_id_ranges.lines() {
        add_fresh_ids(new_id_range, &mut fresh_id_ranges);
    }

    fresh_id_ranges.sort();

    get_total_ids(&fresh_id_ranges)
}

fn get_total_ids(fresh_id_ranges: &Vec<(i64, i64)>) -> i64 {
    let mut total = 0;
    let mut index = 0;

    while index < fresh_id_ranges.len() {
        let number_of_ids;
        let mut index2 = index;
        let mut max_upper = fresh_id_ranges[index].1;

        while index2 + 1 < fresh_id_ranges.len() && fresh_id_ranges[index2 + 1].0 <= max_upper {
            index2 += 1;
            max_upper = max_upper.max(fresh_id_ranges[index2].1);
        }

        number_of_ids = max_upper - fresh_id_ranges[index].0 + 1;
        total += number_of_ids;
        index = index2 + 1;
    }

    total
}

fn add_fresh_ids(new_id_range: &str, fresh_id_ranges: &mut Vec<(i64, i64)>) {
    let ids: Vec<&str> = new_id_range.split('-').collect();

    let num1: i64 = ids.get(0).unwrap().parse().unwrap();
    let num2: i64 = ids.get(1).unwrap().parse().unwrap();

    fresh_id_ranges.push((num1, num2));
}

fn is_fresh(id: &str, fresh_id_ranges: &Vec<(i64, i64)>) -> bool {
    let mut fresh = false;
    let num_id: i64 = id.parse().unwrap();

    for range in fresh_id_ranges {
        if range.0 <= num_id && num_id <= range.1 {
            fresh = true;
            break;
        }
    }

    fresh
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add_fresh_ids() {
        let mut fresh_id_ranges: Vec<(i64, i64)> = Vec::new();

        add_fresh_ids("3-5", &mut fresh_id_ranges);
        add_fresh_ids("10-14", &mut fresh_id_ranges);
        add_fresh_ids("16-20", &mut fresh_id_ranges);
        add_fresh_ids("12-18", &mut fresh_id_ranges);

        assert!(fresh_id_ranges.contains(&(3, 5)));
        assert!(fresh_id_ranges.contains(&(10, 14)));
        assert!(fresh_id_ranges.contains(&(16, 20)));
        assert!(fresh_id_ranges.contains(&(12, 18)));
    }

    #[test]
    fn test_is_fresh() {
        let mut fresh_id_ranges: Vec<(i64, i64)> = Vec::new();

        add_fresh_ids("3-5", &mut fresh_id_ranges);
        add_fresh_ids("10-14", &mut fresh_id_ranges);
        add_fresh_ids("16-20", &mut fresh_id_ranges);
        add_fresh_ids("12-18", &mut fresh_id_ranges);

        assert!(!is_fresh("1", &fresh_id_ranges));
        assert!(is_fresh("5", &fresh_id_ranges));
        assert!(!is_fresh("8", &fresh_id_ranges));
        assert!(is_fresh("11", &fresh_id_ranges));
        assert!(is_fresh("17", &fresh_id_ranges));
        assert!(!is_fresh("32", &fresh_id_ranges));
    }

    #[test]
    fn test_get_total_ids() {
        let mut fresh_id_ranges: Vec<(i64, i64)> = Vec::new();

        add_fresh_ids("3-5", &mut fresh_id_ranges);
        add_fresh_ids("10-14", &mut fresh_id_ranges);
        add_fresh_ids("16-20", &mut fresh_id_ranges);
        add_fresh_ids("12-18", &mut fresh_id_ranges);

        fresh_id_ranges.sort();

        assert_eq!(get_total_ids(&fresh_id_ranges), 14);
    }

    #[test]
    fn test_get_total_ids_max_upper() {
        let mut fresh_id_ranges: Vec<(i64, i64)> = Vec::new();
        add_fresh_ids("10-20", &mut fresh_id_ranges);
        add_fresh_ids("12-13", &mut fresh_id_ranges);
        add_fresh_ids("18-25", &mut fresh_id_ranges);
        fresh_id_ranges.sort();

        assert_eq!(get_total_ids(&fresh_id_ranges), 16);
    }
}
