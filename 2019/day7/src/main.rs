use std::collections::VecDeque;
use std::io;
use std::io::BufRead;

fn main() {
    let stdin = io::stdin();
    let mut numbers = vec![];
    for line in stdin.lock().lines() {
        numbers.extend(line.unwrap()
            .split(",")
            .map(|s| s.parse::<i32>().unwrap())
            .collect::<Vec<i32>>());
    }
    let mut biggest_thrust = 0;
    for permutation in permutations(vec![0, 1, 2, 3, 4]) {
        let mut wire = 0;
        for signal in permutation {
            let mut computer = Computer::new(&numbers);
            computer.inputs.push_back(signal);
            computer.inputs.push_back(wire);
            wire = computer.run();
        }
        if wire > biggest_thrust {
            biggest_thrust = wire;
        }
    }
    println!("{}", biggest_thrust);
    
    let mut feedback_biggest_thrust = 0;
    for permutation in permutations(vec![5, 6, 7, 8, 9]) {
        let mut wire = 0;
        let mut computers = vec![Computer::new(&numbers); 5];
        for i in 0..5 {
            let computer = &mut computers[i];
            computer.inputs.push_back(permutation[i]);
        }
        loop {
            for i in 0..5 {
                let computer = &mut computers[i];
                computer.inputs.push_back(wire);
                wire = computer.run();
                if wire > feedback_biggest_thrust {
                    feedback_biggest_thrust = wire;
                }
            }
            if computers.iter().all(|c| c.halted) {
                break;
            }
        }
    }
    println!("{}", feedback_biggest_thrust);
}

#[derive(Clone)]
struct Computer {
    mem: Vec<i32>,
    modes: Vec<i32>,
    pointer: usize,
    inputs: VecDeque<i32>,
    halted: bool,
}

impl Computer {
    fn new(program: &Vec<i32>) -> Self {
        Computer {
            mem: program.clone(),
            modes: vec![],
            pointer: 0,
            inputs: VecDeque::new(),
            halted: false,
        }
    }

    fn run(&mut self) -> i32 {
        loop {
            let instruction = self.mem[self.pointer];
            let opcode = digits(instruction, 0, 2);
            self.modes = vec![
                digits(instruction, 2, 3),
                digits(instruction, 3, 4),
                digits(instruction, 4, 5),
            ];
            match opcode {
                1 => {
                    let left = self.param(0);
                    let right = self.param(1);
                    let target = self.param(2);
                    self.mem[target] = self.mem[left] + self.mem[right];
                    self.pointer += 4;
                }
                2 => {
                    let left = self.param(0);
                    let right = self.param(1);
                    let target = self.param(2);
                    self.mem[target] = self.mem[left] * self.mem[right];
                    self.pointer += 4;
                }
                3 => {
                    let target = self.param(0);
                    self.mem[target] = self.inputs.pop_front().unwrap();
                    self.pointer += 2;
                }
                4 => {
                    let target = self.param(0);
                    self.pointer += 2;
                    return self.mem[target];
                }
                5 => {
                    let address = self.param(0);
                    let value = self.mem[address];
                    if value != 0 {
                        let target = self.param(1);
                        self.pointer = self.mem[target] as usize;
                    } else {
                        self.pointer += 3
                    }
                }
                6 => {
                    let address = self.param(0);
                    let value = self.mem[address];
                    if value == 0 {
                        let target = self.param(1);
                        self.pointer = self.mem[target] as usize;
                    } else {
                        self.pointer += 3
                    }
                }
                7 => {
                    let left = self.param(0);
                    let right = self.param(1);
                    let target = self.param(2);
                    self.mem[target] = if self.mem[left] < self.mem[right] { 1 } else { 0 };
                    self.pointer += 4;
                }
                8 => {
                    let left = self.param(0);
                    let right = self.param(1);
                    let target = self.param(2);
                    self.mem[target] = if self.mem[left] == self.mem[right] { 1 } else { 0 };
                    self.pointer += 4;
                }
                99 => {
                    self.halted = true;
                    break;
                }
                _ => {}
            }
        }
        return 0;
    }

    fn param(&self, n: usize) -> usize {
        if self.modes[n] == 0 {
            self.mem[self.pointer + n + 1] as usize
        } else {
            self.pointer + n + 1
        }
    }
}


fn digits(n: i32, from: i32, to: i32) -> i32 {
    (n % 10_i32.pow((to) as u32)) / 10_i32.pow(from as u32)
}

fn permutations(elements: Vec<i32>) -> Vec<Vec<i32>> {
    let mut result = vec![];
    if elements.len() <= 1 {
        result.push(elements);
    } else {
        for permutation in permutations(elements[1..].to_vec()) {
            for i in 0..elements.len() {
                let mut v = vec![];
                for e in permutation[..i].to_vec() {
                    v.push(e);
                }
                for e in elements[0..1].to_vec() {
                    v.push(e);
                }
                for e in permutation[i..].to_vec() {
                    v.push(e);
                }
                result.push(v);
            }
        }
    }
    result
}
