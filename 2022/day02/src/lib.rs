use std::str::FromStr;

pub fn part1(input: &str) -> i32 {
    input
        .lines()
        .map(|line| (parse_hand(line, 0), parse_hand(line, 2)))
        .map(|(opponent, me)| me.game(opponent))
        .sum::<i32>()
}

pub fn part2(input: &str) -> i32 {
    input
        .lines()
        .map(|line| (parse_hand(line, 0), parse_outcome(line, 2)))
        .map(|(opponent, desired_outcome)| {
            (
                opponent,
                Hand::for_desired_outcome(opponent, desired_outcome),
            )
        })
        .map(|(opponent, me)| me.game(opponent))
        .sum::<i32>()
}

fn parse_hand(line: &str, pos: usize) -> Hand {
    line.chars()
        .nth(pos)
        .unwrap()
        .to_string()
        .parse::<Hand>()
        .unwrap()
}

fn parse_outcome(line: &str, pos: usize) -> Outcome {
    line.chars()
        .nth(pos)
        .unwrap()
        .to_string()
        .parse::<Outcome>()
        .unwrap()
}

#[derive(Clone, Copy)]
enum Hand {
    Rock,
    Paper,
    Scissors,
}

impl Hand {
    fn score(&self) -> i32 {
        match self {
            Hand::Rock => 1,
            Hand::Paper => 2,
            Hand::Scissors => 3,
        }
    }

    fn game(&self, other: Hand) -> i32 {
        match (self, other) {
            (Hand::Rock, Hand::Scissors)
            | (Hand::Scissors, Hand::Paper)
            | (Hand::Paper, Hand::Rock) => self.score() + 6,
            (Hand::Rock, Hand::Rock)
            | (Hand::Paper, Hand::Paper)
            | (Hand::Scissors, Hand::Scissors) => self.score() + 3,
            (Hand::Scissors, Hand::Rock)
            | (Hand::Paper, Hand::Scissors)
            | (Hand::Rock, Hand::Paper) => self.score(),
        }
    }

    fn for_desired_outcome(opponent: Hand, desired: Outcome) -> Hand {
        match (opponent, desired) {
            (Hand::Rock, Outcome::Loss) => Hand::Scissors,
            (Hand::Paper, Outcome::Loss) => Hand::Rock,
            (Hand::Scissors, Outcome::Loss) => Hand::Paper,
            (Hand::Rock, Outcome::Draw) => Hand::Rock,
            (Hand::Paper, Outcome::Draw) => Hand::Paper,
            (Hand::Scissors, Outcome::Draw) => Hand::Scissors,
            (Hand::Rock, Outcome::Win) => Hand::Paper,
            (Hand::Paper, Outcome::Win) => Hand::Scissors,
            (Hand::Scissors, Outcome::Win) => Hand::Rock,
        }
    }
}

impl FromStr for Hand {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "A" | "X" => Ok(Hand::Rock),
            "B" | "Y" => Ok(Hand::Paper),
            "C" | "Z" => Ok(Hand::Scissors),
            other => Err(format!("invalid hand: {}", other)),
        }
    }
}

enum Outcome {
    Loss,
    Draw,
    Win,
}

impl FromStr for Outcome {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "X" => Ok(Outcome::Loss),
            "Y" => Ok(Outcome::Draw),
            "Z" => Ok(Outcome::Win),
            other => Err(format!("invalid outcome: {}", other)),
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    const INPUT: &str = "A Y
B X
C Z";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 15);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), 12);
    }
}
