#![feature(iter_array_chunks)]

use std::collections::HashSet;

pub fn part1(input: &str) -> i32 {
    input
        .lines()
        .map(|line| line.split_at(line.len() / 2))
        .map(|(left, right)| {
            let a: HashSet<char> = HashSet::from_iter(left.chars());
            let b: HashSet<char> = HashSet::from_iter(right.chars());
            (a, b)
        })
        .map(|(left, right)| *left.intersection(&right).next().unwrap())
        .map(score)
        .sum::<i32>()
}

pub fn part2(input: &str) -> i32 {
    input
        .lines()
        .array_chunks::<3>()
        .map(|rucksacks| {
            rucksacks[0]
                .chars()
                .find(|c| rucksacks[1].contains(*c) && rucksacks[2].contains(*c))
                .unwrap()
        })
        .map(score)
        .sum::<i32>()
}

fn score(c: char) -> i32 {
    if c.is_ascii_lowercase() {
        (c as i32) - 96
    } else if c.is_ascii_uppercase() {
        (c as i32) - 38
    } else {
        panic!("unexpected char {c}")
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";

    #[test]
    fn part_1_works() {
        assert_eq!(part1(INPUT), 157);
    }

    #[test]
    fn part_2_works() {
        assert_eq!(part2(INPUT), 70);
    }
}
