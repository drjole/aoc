use regex::Regex;

pub fn part1(input: &str) -> i32 {
    let check = |a: &[i32], b: &[i32]| a[0] <= b[0] && a[1] >= b[1];
    process(input, check)
}

pub fn part2(input: &str) -> i32 {
    let check = |a: &[i32], b: &[i32]| (a[0]..=a[1]).any(|n| (b[0]..=b[1]).contains(&n));
    process(input, check)
}

fn process(input: &str, check: fn(&[i32], &[i32]) -> bool) -> i32 {
    let pattern = Regex::new(r"(\d+)-(\d+),(\d+)-(\d+)").unwrap();
    input
        .lines()
        .map(|line| {
            let numbers = pattern
                .captures(line)
                .unwrap()
                .iter()
                .skip(1)
                .map(Option::unwrap)
                .map(|c| c.as_str().parse::<i32>().unwrap())
                .collect::<Vec<_>>();
            i32::from(check(&numbers[..2], &numbers[2..]) || check(&numbers[2..], &numbers[..2]))
        })
        .sum::<i32>()
}

#[cfg(test)]
mod test {
    use super::*;

    const INPUT: &str = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 2);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), 4);
    }
}
