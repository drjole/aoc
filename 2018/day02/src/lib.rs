use std::collections::HashMap;

pub fn part1(input: &str) -> i32 {
    let mut all_counts: Vec<HashMap<char, i32>> = vec![];
    input.lines().for_each(|line| {
        let mut counts: HashMap<char, i32> = HashMap::new();
        line.chars().for_each(|c| {
            counts.entry(c).and_modify(|e| *e += 1).or_insert(1);
        });
        all_counts.push(counts);
    });
    let mut result = all_counts
        .iter()
        .filter(|counts| counts.iter().any(|(_key, &value)| value == 2))
        .count() as i32;
    result *= all_counts
        .iter()
        .filter(|counts| counts.iter().any(|(_key, &value)| value == 3))
        .count() as i32;
    result
}

pub fn part2(_input: &str) -> i32 {
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 12);
    }
}
