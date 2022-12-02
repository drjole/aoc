pub fn part1(input: &str) -> i32 {
    preprocess(input).into_iter().max().unwrap()
}

pub fn part2(input: &str) -> i32 {
    let mut result = preprocess(input);
    result.sort_by(|a, b| b.cmp(a));
    result.iter().take(3).sum::<i32>()
}

fn preprocess(input: &str) -> Vec<i32> {
    input
        .split("\n\n")
        .map(|payload| {
            payload
                .lines()
                .map(|value| value.parse::<i32>().unwrap())
                .sum::<i32>()
        })
        .collect::<Vec<_>>()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 24000);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), 45000);
    }
}
