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
    println!("{}", execute(&numbers, 1));
    println!("{}", execute(&numbers, 5));
}

fn execute(program: &Vec<i32>, input: i32) -> i32 {
    let mut memory = program.clone();
    let mut pointer = 0;
    let mut output = 0;
    loop {
        let instruction = memory[pointer];
        let opcode = digits(instruction, 0, 2);
        let modes = vec![
            digits(instruction, 2, 3),
            digits(instruction, 3, 4),
            digits(instruction, 4, 5),
        ];
        match opcode {
            1 => {
                let left = param(&memory, &modes, pointer, 0);
                let right = param(&memory, &modes, pointer, 1);
                let target = param(&memory, &modes, pointer, 2);
                memory[target] = memory[left] + memory[right];
                pointer += 4;
            }
            2 => {
                let left = param(&memory, &modes, pointer, 0);
                let right = param(&memory, &modes, pointer, 1);
                let target = param(&memory, &modes, pointer, 2);
                memory[target] = memory[left] * memory[right];
                pointer += 4;
            }
            3 => {
                let target = param(&memory, &modes, pointer, 0);
                memory[target] = input;
                pointer += 2;
            }
            4 => {
                let target = param(&memory, &modes, pointer, 0);
                output = memory[target];
                pointer += 2;
            }
            5 => {
                let address = param(&memory, &modes, pointer, 0);
                let value = memory[address];
                if value != 0 {
                    let target = param(&memory, &modes, pointer, 1);
                    pointer = memory[target] as usize;
                } else {
                    pointer += 3
                }
            }
            6 => {
                let address = param(&memory, &modes, pointer, 0);
                let value = memory[address];
                if value == 0 {
                    let target = param(&memory, &modes, pointer, 1);
                    pointer = memory[target] as usize;
                } else {
                    pointer += 3
                }
            }
            7 => {
                let left = param(&memory, &modes, pointer, 0);
                let right = param(&memory, &modes, pointer, 1);
                let target = param(&memory, &modes, pointer, 2);
                memory[target] = if memory[left] < memory[right] { 1 } else { 0 };
                pointer += 4;
            }
            8 => {
                let left = param(&memory, &modes, pointer, 0);
                let right = param(&memory, &modes, pointer, 1);
                let target = param(&memory, &modes, pointer, 2);
                memory[target] = if memory[left] == memory[right] { 1 } else { 0 };
                pointer += 4;
            }
            99 => break,
            _ => {}
        }
    }
    return output;
}

fn param(memory: &Vec<i32>, modes: &Vec<i32>, pointer: usize, n: usize) -> usize {
    if modes[n] == 0 {
        memory[pointer + n + 1] as usize
    } else {
        pointer + n + 1
    }
}

fn digits(n: i32, from: i32, to: i32) -> i32 {
    (n % 10_i32.pow((to) as u32)) / 10_i32.pow(from as u32)
}
