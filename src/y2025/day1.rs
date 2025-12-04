use std::fs::read_to_string;

pub fn part1() -> i32 {
    let mut dial = 50;
    let mut password = 0;

    for instruction in read_to_string("./inputs/2025day1.txt").unwrap().lines() {
        dial = move_dial(dial, instruction);
        // println!("dial: {dial}, instruction: {instruction}, password: {password}");
        if dial == 0 {
            password += 1;
        }
    }

    return password;
}

pub fn part2() -> i32 {
    let mut dial = 50;
    let mut password = 0;

    for instruction in read_to_string("./inputs/2025day1.txt").unwrap().lines() {
        password += move_dial_2(dial, instruction);
        dial = move_dial(dial, instruction);
        // println!("dial: {dial}, instruction: {instruction}, password: {password}");
    }

    return password;
}

fn move_dial(dial: i32, instruction: &str) -> i32 {
    let rotate = &instruction[1..];
    let rotate_num: i32 = rotate.parse().unwrap();

    match &instruction[0..1] {
        "L" => ((dial - rotate_num) % 100 + 100) % 100,
        "R" => (dial + rotate_num) % 100,
        _ => panic!("invalid input"),
    }
}

fn move_dial_2(dial: i32, instruction: &str) -> i32 {
    let rotate = &instruction[1..];
    let rotate_num: i32 = rotate.parse().unwrap();

    match &instruction[0..1] {
        "L" => (((100 - dial) % 100) + rotate_num) / 100,
        "R" => (dial + rotate_num) / 100,
        _ => panic!("invalid input"),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_move_dial() {
        assert_eq!(move_dial(50, "L32"), 18);
        assert_eq!(move_dial(50, "R32"), 82);
    }

    #[test]
    fn test_move_dial_overflow() {
        assert_eq!(move_dial(50, "L55"), 95);
        assert_eq!(move_dial(50, "R87"), 37);
        assert_eq!(move_dial(50, "R320"), 70)
    }

    #[test]
    fn test_move_dial_multiple() {
        let instructions = [
            "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
        ];
        let mut dial = 50;

        for instruction in instructions {
            dial = move_dial(dial, instruction);
            println!("{dial}");
        }

        assert_eq!(dial, 32);
    }

    #[test]
    fn test_password() {
        let instructions = [
            "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
        ];
        let mut dial = 50;
        let mut password = 0;

        for instruction in instructions {
            dial = move_dial(dial, instruction);
            if dial == 0 {
                password += 1;
            }
        }

        assert_eq!(password, 3);
    }

    #[test]
    fn test_move_dial_2() {
        assert_eq!(move_dial_2(50, "L32"), 0);
        assert_eq!(move_dial_2(50, "R32"), 0);
        assert_eq!(move_dial_2(50, "L50"), 1);
    }

    #[test]
    fn test_move_dial_2_overflow() {
        assert_eq!(move_dial_2(50, "L55"), 1);
        assert_eq!(move_dial_2(50, "R87"), 1);
        assert_eq!(move_dial_2(50, "L68"), 1);
        assert_eq!(move_dial_2(50, "R320"), 3);
        assert_eq!(move_dial_2(50, "L320"), 3);
        assert_eq!(move_dial_2(10, "L20"), 1);
    }

    #[test]
    fn test_password_2() {
        let instructions = [
            "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
        ];
        let dial_positions = [82, 52, 0, 95, 55, 0, 99, 0, 14, 32];

        let mut dial = 50;
        let mut password = 0;

        for (index, instruction) in instructions.iter().enumerate() {
            password += move_dial_2(dial, instruction);
            dial = move_dial(dial, instruction);
            // println!("dial: {dial}, instruction: {instruction}, password: {password}");
            assert_eq!(dial, dial_positions[index]);
        }

        assert_eq!(password, 6);
    }
}
