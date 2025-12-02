use std::fs::read_to_string;

pub fn part1() -> i64 {
    let database = read_to_string("./inputs/day2.txt").unwrap();
    let ranges = get_product_ranges(&database.trim());
    let mut sum: i64 = 0;

    for range in ranges {
        let ids = get_ids(range);

        let id_range = generate_range(ids[0], ids[1]);

        for id in id_range {
            if invalid_id(&id) {
                let invalid_id: i64 = id.parse().unwrap();
                sum += invalid_id;
            }
        }
    }

    return sum;
}

fn get_product_ranges(database: &str) -> Vec<&str> {
    database.split_terminator(",").collect()
}

fn get_ids(range: &str) -> Vec<&str> {
    range.split_terminator("-").collect()
}

fn generate_range(first_id: &str, last_id: &str) -> Vec<String> {
    let mut ids: Vec<String> = Vec::new();

    let mut first_id_num: i64 = first_id.parse().unwrap();
    let last_id_num: i64 = last_id.parse().unwrap();

    while first_id_num <= last_id_num {
        ids.push(first_id_num.to_string());
        first_id_num += 1;
    }

    ids
}

fn invalid_id(id: &str) -> bool {
    let ptr1 = 0;
    let ptr2 = id.len() / 2;
    let mut invalid_id = true;

    if id[ptr1..ptr2] != id[ptr2..] {
        invalid_id = false;
    }

    invalid_id
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_invalid_id_true() {
        assert!(invalid_id("11"));
        assert!(invalid_id("22"));
        assert!(invalid_id("99"));
        assert!(invalid_id("1188511885"));
        assert!(invalid_id("222222"));
        assert!(invalid_id("446446"));
        assert!(invalid_id("38593859"));
        assert!(invalid_id("122122"));
    }

    #[test]
    fn test_invalid_id_false() {
        assert!(!invalid_id("12"));
        assert!(!invalid_id("23"));
        assert!(!invalid_id("98"));
        assert!(!invalid_id("1188411885"));
        assert!(!invalid_id("222122"));
        assert!(!invalid_id("247486"));
        assert!(!invalid_id("38393833"));
    }

    #[test]
    fn test_generate_range() {
        assert_eq!(generate_range("11", "13"), vec!["11", "12", "13"]);
        assert_eq!(generate_range("1698522", "1698528").len(), 7);
    }

    #[test]
    fn test_get_ids() {
        assert_eq!(get_ids("11-22"), vec!["11", "22"]);
        assert_eq!(
            get_ids("1188511880-1188511890"),
            vec!["1188511880", "1188511890"]
        );
    }

    #[test]
    fn test_get_product_ranges() {
        let database = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
        let ranges: Vec<&str> = database.split(',').map(|s| s.trim()).collect();

        assert_eq!(get_product_ranges(database), ranges);
    }
}
