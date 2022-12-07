use std::collections::HashSet;

pub fn part1(input: &str) -> i32 {
    let mut frequency = 0;
    input
        .lines()
        .map(|line| line.parse::<i32>().unwrap())
        .for_each(|n| frequency += n);
    frequency
}

pub fn part2(input: &str) -> i32 {
    let mut frequency = 0;
    let mut frequencies = HashSet::from([frequency]);
    let _ = input
        .lines()
        .cycle()
        .map(|line| line.parse::<i32>().unwrap())
        .map_while(|n| {
            frequency += n;
            if frequencies.contains(&frequency) {
                None
            } else {
                frequencies.insert(frequency);
                Some(())
            }
        })
        .collect::<Vec<_>>();
    frequency
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_works() {
        assert_eq!(part1(prepare_input("+1, +1, +1").as_str()), 3);
        assert_eq!(part1(prepare_input("+1, +1, -2").as_str()), 0);
        assert_eq!(part1(prepare_input("-1, -2, -3").as_str()), -6);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(prepare_input("+1, -1").as_str()), 0);
        assert_eq!(part2(prepare_input("+3, +3, +4, -2, -4").as_str()), 10);
        assert_eq!(part2(prepare_input("-6, +3, +8, +5, -6").as_str()), 5);
        assert_eq!(part2(prepare_input("+7, +7, -2, -7, -4").as_str()), 14);
    }

    fn prepare_input(input: &str) -> String {
        input.replace(',', "").replace(' ', "\n")
    }
}
