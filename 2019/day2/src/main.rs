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
    println!("{}", execute(&numbers, 12, 2));
    for i in 0..100 {
        for j in 0..100 {
            if execute(&numbers, i, j) == 19690720 {
                println!("{}", 100 * i + j);
            }
        }
    }
}

fn execute(program: &Vec<i32>, a: i32, b: i32) -> i32 {
    let mut memory = program.clone();
    memory[1] = a;
    memory[2] = b;
    let mut pointer = 0;
    loop {
        let left = memory[pointer + 1] as usize;
        let right = memory[pointer + 2] as usize;
        let target_address = memory[pointer + 3] as usize;
        match memory[pointer] {
            1 => memory[target_address] = memory[left] + memory[right],
            2 => memory[target_address] = memory[left] * memory[right],
            99 => break,
            _ => {}
        }
        pointer += 4;
    }
    return memory[0];
}
