use nom::branch::alt;
use nom::bytes::complete::tag;
use nom::character::complete::newline;
use nom::multi::separated_list1;
use nom::IResult;

pub fn part1(input: &str) -> i32 {
    let cpu = process(input);
    cpu.captured_signal_strengths.iter().sum::<i32>()
}

pub fn part2(input: &str) -> String {
    let cpu = process(input);
    cpu.screen
}

fn process(input: &str) -> Cpu {
    let instructions = parse(input).unwrap().1;
    let mut cpu = Cpu::new();
    cpu.capture_signal_strengths(vec![20, 60, 100, 140, 180, 220]);
    for instruction in instructions {
        cpu.execute(instruction);
    }
    cpu
}

struct Cpu {
    clock: u32,
    x: i32,
    capture_cycles: Vec<u32>,
    captured_signal_strengths: Vec<i32>,
    screen: String,
    screen_position: i32,
    screen_row: i32,
}

impl Cpu {
    fn new() -> Self {
        Cpu {
            clock: 0,
            x: 1,
            capture_cycles: vec![],
            captured_signal_strengths: vec![],
            screen: "".to_string(),
            screen_position: 0,
            screen_row: 0,
        }
    }
    fn capture_signal_strengths(&mut self, cycles: Vec<u32>) -> &Self {
        self.capture_cycles = cycles;
        self
    }

    fn execute(&mut self, instruction: Instruction) {
        for _ in 0..instruction.cycles() {
            self.clock += 1;
            if self.capture_cycles.contains(&self.clock) {
                self.captured_signal_strengths
                    .push(self.clock as i32 * self.x);
            }
            let pixel = if (self.screen_position - 1..=self.screen_position + 1).contains(&self.x) {
                '#'
            } else {
                '.'
            };
            self.screen.push(pixel);
            self.screen_position += 1;
            if self.screen_position == 40 {
                self.screen_position = 0;
                self.screen_row += 1;
                if self.screen_row != 6 {
                    self.screen.push('\n');
                }
            }
        }
        match instruction {
            Instruction::AddX(n) => self.x += n,
            Instruction::Noop => {}
        }
    }
}

#[derive(Debug)]
enum Instruction {
    AddX(i32),
    Noop,
}

impl Instruction {
    fn cycles(&self) -> u32 {
        match self {
            Instruction::AddX(_) => 2,
            Instruction::Noop => 1,
        }
    }
}

fn parse(input: &str) -> IResult<&str, Vec<Instruction>> {
    separated_list1(newline, alt((addx, noop)))(input)
}

fn addx(input: &str) -> IResult<&str, Instruction> {
    let (input, _) = tag("addx ")(input)?;
    let (input, n) = nom::character::complete::i32(input)?;
    let instruction = Instruction::AddX(n);
    Ok((input, instruction))
}

fn noop(input: &str) -> IResult<&str, Instruction> {
    let (input, _) = tag("noop")(input)?;
    let instruction = Instruction::Noop;
    Ok((input, instruction))
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 13140)
    }

    #[test]
    fn part2_works() {
        assert_eq!(
            part2(INPUT),
            "##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######....."
        )
    }
}
