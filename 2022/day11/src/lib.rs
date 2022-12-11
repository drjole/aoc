use std::collections::VecDeque;

use nom::branch::alt;
use nom::bytes::complete::tag;
use nom::character::complete;
use nom::character::complete::multispace1;
use nom::multi::separated_list1;
use nom::IResult;

pub fn part1(input: &str) -> i64 {
    process(input, 20, false)
}

pub fn part2(input: &str) -> i64 {
    process(input, 10000, true)
}

fn process(input: &str, rounds: i64, crazy_worried: bool) -> i64 {
    let mut monkeys = parse_monkeys(input).unwrap().1;
    if crazy_worried {
        monkeys
            .iter_mut()
            .for_each(|monkey| monkey.crazy_worried = true);
    }
    let magic_number = monkeys.iter().map(|monkey| monkey.divisor).product::<i64>();
    for _ in 0..rounds {
        for monkey_index in 0..monkeys.len() {
            for _ in 0..monkeys[monkey_index].items.len() {
                let monkey = monkeys.get_mut(monkey_index).unwrap();
                monkey.business += 1;
                let (catcher_id, item) = monkey.throw(magic_number);
                let catcher = monkeys.get_mut(catcher_id).unwrap();
                catcher.catch(item);
            }
        }
    }
    let mut businesses = monkeys
        .iter()
        .map(|monkey| monkey.business)
        .collect::<Vec<_>>();
    businesses.sort_by(|a, b| b.cmp(a));
    businesses[0] * businesses[1]
}

#[derive(Debug)]
struct Monkey {
    items: VecDeque<i64>,
    operation: Operation,
    divisor: i64,
    throw_true: usize,
    throw_false: usize,
    business: i64,
    crazy_worried: bool,
}

impl Monkey {
    fn throw(&mut self, magic_number: i64) -> (usize, i64) {
        let item = self.items.pop_front().unwrap();
        let left = match self.operation.left {
            OperationInput::Old => item,
            OperationInput::Number(n) => n,
        };
        let right = match self.operation.right {
            OperationInput::Old => item,
            OperationInput::Number(n) => n,
        };
        let mut item = match self.operation.operator {
            Operator::Add => left + right,
            Operator::Multiply => (left * right) % magic_number,
        };
        if !self.crazy_worried {
            item /= 3;
        }
        let catcher_id = if item % self.divisor == 0 {
            self.throw_true
        } else {
            self.throw_false
        };
        (catcher_id, item)
    }

    fn catch(&mut self, item: i64) {
        self.items.push_back(item);
    }
}

#[derive(Debug)]
struct Operation {
    left: OperationInput,
    operator: Operator,
    right: OperationInput,
}

#[derive(Debug)]
enum OperationInput {
    Old,
    Number(i64),
}

#[derive(Debug)]
enum Operator {
    Add,
    Multiply,
}

fn parse_monkeys(input: &str) -> IResult<&str, Vec<Monkey>> {
    separated_list1(multispace1, parse_monkey)(input)
}

fn parse_monkey(input: &str) -> IResult<&str, Monkey> {
    let (input, _) = tag("Monkey ")(input)?;
    let (input, _) = complete::i64(input)?;
    let (input, _) = tag(":")(input)?;
    let (input, _) = complete::newline(input)?;
    let (input, _) = tag("  Starting items: ")(input)?;
    let (input, items) = separated_list1(tag(", "), complete::i64)(input)?;
    let (input, _) = complete::newline(input)?;
    let (input, _) = tag("  Operation: new = ")(input)?;
    let (input, operation) = parse_operation(input)?;
    let (input, _) = complete::newline(input)?;
    let (input, _) = tag("  Test: divisible by ")(input)?;
    let (input, divisor) = complete::i64(input)?;
    let (input, _) = complete::newline(input)?;
    let (input, _) = tag("    If true: throw to monkey ")(input)?;
    let (input, throw_true) = complete::i64(input)?;
    let (input, _) = complete::newline(input)?;
    let (input, _) = tag("    If false: throw to monkey ")(input)?;
    let (input, throw_false) = complete::i64(input)?;
    let monkey = Monkey {
        items: VecDeque::from(items),
        operation,
        divisor,
        throw_true: throw_true as usize,
        throw_false: throw_false as usize,
        business: 0,
        crazy_worried: false,
    };
    Ok((input, monkey))
}

fn parse_operation(input: &str) -> IResult<&str, Operation> {
    let (input, left) = parse_operation_input(input)?;
    let (input, _) = tag(" ")(input)?;
    let (input, operator) = parse_operator(input)?;
    let (input, _) = tag(" ")(input)?;
    let (input, right) = parse_operation_input(input)?;
    let operation = Operation {
        left,
        operator,
        right,
    };
    Ok((input, operation))
}

fn parse_operation_input(input: &str) -> IResult<&str, OperationInput> {
    let (input, operation_input) = alt((tag("old"), complete::digit1))(input)?;
    let operation_input = match operation_input {
        "old" => OperationInput::Old,
        other => OperationInput::Number(other.parse::<i64>().unwrap()),
    };
    Ok((input, operation_input))
}

fn parse_operator(input: &str) -> IResult<&str, Operator> {
    let (input, operator) = alt((tag("*"), tag("+")))(input)?;
    let operator = match operator {
        "*" => Operator::Multiply,
        "+" => Operator::Add,
        other => panic!("Invalid operator: {}", other),
    };
    Ok((input, operator))
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 10605);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), 2713310158);
    }
}
