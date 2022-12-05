use regex::Regex;

pub fn part1(input: &str) -> String {
    process(input, |stacks, count, from, to| {
        for _ in 0..count {
            let v = stacks[from as usize].pop().unwrap();
            stacks[to as usize].push(v);
        }
    })
}

pub fn part2(input: &str) -> String {
    process(input, |stacks, count, from, to| {
        let mut payload: Vec<String> = vec![];
        for _ in 0..count {
            payload.push(stacks[from as usize].pop().unwrap());
        }
        for i in (0..count).rev() {
            stacks[to as usize].push(payload[i as usize].clone());
        }
    })
}

fn process(input: &str, execute: fn(&mut Vec<Vec<String>>, i32, i32, i32)) -> String {
    let (initial_stacks, instructions) = input.split_once("\n\n").unwrap();
    let stack_count = (initial_stacks.lines().next().unwrap().len() + 1) / 4;
    let mut stacks: Vec<Vec<String>> = vec![vec![]; stack_count];
    initial_stacks
        .lines()
        .rev()
        .skip(1)
        .flat_map(|line| {
            (0..stack_count)
                .into_iter()
                .map(|n| line.chars().nth(4 * n + 1).unwrap())
        })
        .enumerate()
        .for_each(|(i, c)| {
            if c != ' ' {
                stacks[i % stack_count].push(c.to_string())
            }
        });
    let instruction_pattern = Regex::new(r"move (\d+) from (\d+) to (\d+)").unwrap();
    instructions.lines().for_each(|line| {
        let instruction = instruction_pattern
            .captures(line)
            .unwrap()
            .iter()
            .skip(1)
            .map(Option::unwrap)
            .map(|c| c.as_str().parse::<i32>().unwrap())
            .collect::<Vec<_>>();
        let (count, from, to) = (instruction[0], instruction[1] - 1, instruction[2] - 1);
        execute(&mut stacks, count, from, to);
    });
    stacks
        .iter()
        .map(|stack| stack[stack.len() - 1].clone())
        .collect::<String>()
}

#[cfg(test)]
mod test {
    use super::*;

    const INPUT: &str = "    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), "CMZ");
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), "MCD");
    }
}
