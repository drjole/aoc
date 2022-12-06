use itertools::Itertools;

pub fn part1(input: &str) -> i32 {
    process(input, 4)
}

pub fn part2(input: &str) -> i32 {
    process(input, 14)
}

fn process(input: &str, n: usize) -> i32 {
    let chars = input.chars().collect::<Vec<char>>();
    chars
        .iter()
        .enumerate()
        .skip(n)
        .find(|&(i, _)| {
            (0..n)
                .rev()
                .combinations(2)
                .all(|v| chars[i - v[0]] != chars[i - v[1]])
        })
        .map(|(i, _)| i + 1)
        .unwrap() as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_works() {
        assert_eq!(part1("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 7);
        assert_eq!(part1("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5);
        assert_eq!(part1("nppdvjthqldpwncqszvftbrmjlhg"), 6);
        assert_eq!(part1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10);
        assert_eq!(part1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 19);
        assert_eq!(part2("bvwbjplbgvbhsrlpgdmjqwftvncz"), 23);
        assert_eq!(part2("nppdvjthqldpwncqszvftbrmjlhg"), 23);
        assert_eq!(part2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 29);
        assert_eq!(part2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 26);
    }
}
